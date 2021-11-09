// Copyright (C) 2017 ScyllaDB

package backup

import (
	_ "embed"
	"fmt"

	"github.com/go-openapi/strfmt"
	"github.com/scylladb/scylla-manager/pkg/command/flag"
	"github.com/scylladb/scylla-manager/pkg/managerclient"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

//go:embed validate.yaml
var validateRes []byte

type validateCommand struct {
	*cobra.Command
	client *managerclient.Client

	cluster             string
	location            []string
	deleteOrphanedFiles bool
	parallel            int
	interval            flag.Duration
	startDate           flag.Time
	numRetries          int
}

func NewValidateCommand(client *managerclient.Client) *cobra.Command {
	c := &cobra.Command{
		Use: "validate",
	}
	if err := yaml.Unmarshal(validateRes, c); err != nil {
		panic(err)
	}

	cmd := &validateCommand{
		Command: c,
		client:  client,
	}
	cmd.init()

	c.RunE = func(_ *cobra.Command, args []string) error {
		return cmd.run()
	}
	return c
}

func (cmd *validateCommand) init() {
	w := flag.Wrap(cmd.Flags())
	w.Cluster(&cmd.cluster)
	w.Location(&cmd.location)
	w.Unwrap().BoolVar(&cmd.deleteOrphanedFiles, "delete-orphaned-files", false, "If set data files not belonging to any snapshot will be deleted after the validation.")
	w.Unwrap().IntVar(&cmd.parallel, "parallel", 0, "Number of hosts to analyze in parallel.")

	w.Interval(&cmd.interval)
	w.StartDate(&cmd.startDate)
	w.NumRetries(&cmd.numRetries, cmd.numRetries)
}

func (cmd *validateCommand) run() error {
	t := &managerclient.Task{
		Type:       "validate_backup",
		Enabled:    true,
		Schedule:   cmd.schedule(),
		Properties: make(map[string]interface{}),
	}

	props := t.Properties.(map[string]interface{})
	if len(cmd.location) != 0 {
		props["location"] = cmd.location
	}
	if cmd.deleteOrphanedFiles {
		props["delete_orphaned_files"] = true
	}
	if cmd.parallel > 0 {
		props["parallel"] = cmd.parallel
	}

	id, err := cmd.client.CreateTask(cmd.Context(), cmd.cluster, t)
	if err != nil {
		return err
	}

	fmt.Fprintln(cmd.OutOrStdout(), managerclient.TaskJoin(t.Type, id))

	return nil
}

func (cmd *validateCommand) schedule() *managerclient.Schedule {
	return &managerclient.Schedule{
		Interval:   cmd.interval.String(),
		StartDate:  strfmt.DateTime(cmd.startDate.Time),
		NumRetries: int64(cmd.numRetries),
	}
}
