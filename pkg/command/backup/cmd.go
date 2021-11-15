// Copyright (C) 2017 ScyllaDB

package backup

import (
	_ "embed"
	"fmt"
	"time"

	"github.com/go-openapi/strfmt"
	"github.com/scylladb/scylla-manager/pkg/command/flag"
	"github.com/scylladb/scylla-manager/pkg/managerclient"
	"github.com/scylladb/scylla-manager/pkg/service/scheduler"
	"github.com/spf13/cobra"
	"go.uber.org/atomic"
	"gopkg.in/yaml.v2"
)

//go:embed res.yaml
var res []byte

//go:embed update-res.yaml
var updateRes []byte

type command struct {
	cobra.Command
	client *managerclient.Client
	update bool

	cluster          string
	dc               []string
	location         []string
	keyspace         []string
	retention        int
	rateLimit        []string
	snapshotParallel []string
	uploadParallel   []string
	dryRun           bool
	showTables       bool
	purgeOnly        bool
	enabled          bool
	interval         flag.Duration
	startDate        flag.Time
	numRetries       int
}

func NewCommand(client *managerclient.Client) *cobra.Command {
	updateCmd := newCommand(client, updateRes)
	updateCmd.update = true
	updateCmd.Command.Args = cobra.ExactArgs(1)

	cmd := newCommand(client, res)
	cmd.AddCommand(&updateCmd.Command)

	return &cmd.Command
}

func newCommand(client *managerclient.Client, res []byte) *command {
	cmd := &command{
		client: client,
	}
	if err := yaml.Unmarshal(res, &cmd.Command); err != nil {
		panic(err)
	}
	cmd.init()
	cmd.RunE = func(_ *cobra.Command, args []string) error {
		return cmd.run(args)
	}
	return cmd
}

func (cmd *command) init() {
	defer flag.MustSetUsages(&cmd.Command, res, "location")

	w := flag.Wrap(cmd.Flags())
	w.Cluster(&cmd.cluster)
	w.Location(&cmd.location)
	w.Datacenter(&cmd.dc)
	w.Keyspace(&cmd.keyspace)
	w.Unwrap().IntVar(&cmd.retention, "retention", 7, "")
	w.Unwrap().StringSliceVar(&cmd.rateLimit, "rate-limit", nil, "")
	w.Unwrap().StringSliceVar(&cmd.snapshotParallel, "snapshot-parallel", nil, "")
	w.Unwrap().StringSliceVar(&cmd.uploadParallel, "upload-parallel", nil, "")
	w.Unwrap().BoolVar(&cmd.dryRun, "dry-run", false, "")
	w.Unwrap().BoolVar(&cmd.showTables, "show-tables", false, "")
	w.Unwrap().BoolVar(&cmd.purgeOnly, "purge-only", false, "")

	w.Enabled(&cmd.enabled)
	w.Interval(&cmd.interval)
	w.StartDate(&cmd.startDate)
	w.NumRetries(&cmd.numRetries, cmd.numRetries)
}

func (cmd *command) run(args []string) error {
	var task *managerclient.Task

	if cmd.update {
		taskType, taskID, err := managerclient.TaskSplit(args[0])
		if err != nil {
			return err
		}
		if scheduler.TaskType(taskType) != scheduler.BackupTask {
			return fmt.Errorf("backup update can't handle %s task", taskType)
		}
		task, err = cmd.client.GetTask(cmd.Context(), cmd.cluster, taskType, taskID)
		if err != nil {
			return err
		}
	} else {
		task = &managerclient.Task{
			Type:       "backup",
			Enabled:    cmd.enabled,
			Schedule:   cmd.schedule(),
			Properties: make(map[string]interface{}),
		}
	}

	props := task.Properties.(map[string]interface{})
	if len(cmd.location) > 0 {
		props["location"] = cmd.location
	}
	if len(cmd.dc) > 0 {
		props["dc"] = cmd.dc
	}
	if len(cmd.keyspace) > 0 {
		props["keyspace"] = cmd.keyspace
	}
	if cmd.Flag("retention").Changed {
		props["retention"] = cmd.retention
	}
	if len(cmd.rateLimit) > 0 {
		props["rate_limit"] = cmd.rateLimit
	}
	if len(cmd.snapshotParallel) > 0 {
		props["snapshot_parallel"] = cmd.snapshotParallel
	}
	if len(cmd.uploadParallel) > 0 {
		props["upload_parallel"] = cmd.uploadParallel
	}
	if cmd.Flag("purge-only").Changed {
		props["purge_only"] = cmd.purgeOnly
	}

	if cmd.dryRun {
		stillWaiting := atomic.NewBool(true)
		time.AfterFunc(5*time.Second, func() {
			if stillWaiting.Load() {
				fmt.Fprintf(cmd.OutOrStderr(), "NOTICE: this may take a while, we are performing disk size calculations on the nodes\n")
			}
		})

		res, err := cmd.client.GetBackupTarget(cmd.Context(), cmd.cluster, task)
		stillWaiting.Store(false)
		if err != nil {
			return err
		}

		fmt.Fprintf(cmd.OutOrStderr(), "NOTICE: dry run mode, backup is not scheduled\n\n")
		if cmd.showTables {
			res.ShowTables = -1
		}
		return res.Render(cmd.OutOrStdout())
	}

	if task.ID == "" {
		id, err := cmd.client.CreateTask(cmd.Context(), cmd.cluster, task)
		if err != nil {
			return err
		}
		task.ID = id.String()
	} else if err := cmd.client.UpdateTask(cmd.Context(), cmd.cluster, task); err != nil {
		return err
	}

	fmt.Fprintln(cmd.OutOrStdout(), managerclient.TaskJoin(task.Type, task.ID))
	return nil
}

func (cmd *command) schedule() *managerclient.Schedule {
	return &managerclient.Schedule{
		Interval:   cmd.interval.String(),
		StartDate:  strfmt.DateTime(cmd.startDate.Time),
		NumRetries: int64(cmd.numRetries),
	}
}
