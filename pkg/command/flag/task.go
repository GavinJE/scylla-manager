package flag

import (
	"github.com/go-openapi/strfmt"
	"github.com/scylladb/scylla-manager/pkg/managerclient"
	"github.com/spf13/cobra"
)

type TaskBase struct {
	cobra.Command
	update bool

	enabled    bool
	interval   Duration
	startDate  Time
	numRetries int
}

func NewTaskBase() TaskBase {
	cmd := TaskBase{}
	cmd.init()
	return cmd
}

func NewUpdateTaskBase() TaskBase {
	cmd := TaskBase{
		Command: cobra.Command{
			Args: cobra.ExactArgs(1),
		},
		update: true,
	}
	cmd.init()
	return cmd
}

func (cmd *TaskBase) init() {
	w := Wrap(cmd.Flags())
	w.enabled(&cmd.enabled)
	w.interval(&cmd.interval)
	w.startDate(&cmd.startDate)
	w.numRetries(&cmd.numRetries, cmd.numRetries)
}

func (cmd *TaskBase) Enabled() bool {
	return cmd.enabled
}

func (cmd *TaskBase) Schedule() *managerclient.Schedule {
	return &managerclient.Schedule{
		Interval:   cmd.interval.String(),
		StartDate:  strfmt.DateTime(cmd.startDate.Time),
		NumRetries: int64(cmd.numRetries),
	}
}

func (cmd *TaskBase) Update() bool {
	return cmd.update
}

func (cmd *TaskBase) UpdateTask(task *managerclient.Task) {
	if cmd.Flag("enabled").Changed {
		task.Enabled = cmd.enabled
	}
	if cmd.Flag("interval").Changed {
		task.Schedule.Interval = cmd.interval.String()
	}
	if cmd.Flag("start-date").Changed {
		task.Schedule.StartDate = strfmt.DateTime(cmd.startDate.Time)
	}
	if cmd.Flag("num-retries").Changed {
		task.Schedule.NumRetries = int64(cmd.numRetries)
	}
}
