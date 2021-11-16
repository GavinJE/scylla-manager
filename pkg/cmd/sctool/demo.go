// Copyright (C) 2017 ScyllaDB

package main

import (
	"github.com/scylladb/scylla-manager/pkg/command/backup"
	"github.com/scylladb/scylla-manager/pkg/command/backup/backupdelete"
	"github.com/scylladb/scylla-manager/pkg/command/backup/backupfiles"
	"github.com/scylladb/scylla-manager/pkg/command/backup/backuplist"
	"github.com/scylladb/scylla-manager/pkg/command/backup/backupvalidate"
	"github.com/scylladb/scylla-manager/pkg/command/repair"
	"github.com/scylladb/scylla-manager/pkg/command/repair/repaircontrol"
	"github.com/scylladb/scylla-manager/pkg/command/resume"
	"github.com/scylladb/scylla-manager/pkg/command/status"
	"github.com/scylladb/scylla-manager/pkg/command/suspend"
	"github.com/spf13/cobra"
)

func init() {
	n := &cobra.Command{
		Use: "new",
	}
	n.AddCommand(backupvalidate.NewCommand(&client))
	n.AddCommand(backuplist.NewCommand(&client))
	n.AddCommand(backupfiles.NewCommand(&client))
	n.AddCommand(backupdelete.NewCommand(&client))
	n.AddCommand(status.NewCommand(&client))
	n.AddCommand(suspend.NewCommand(&client))
	n.AddCommand(resume.NewCommand(&client))
	n.AddCommand(backup.NewCommand(&client))
	n.AddCommand(repair.NewCommand(&client))
	n.AddCommand(repaircontrol.NewCommand(&client))
	rootCmd.AddCommand(n)
}
