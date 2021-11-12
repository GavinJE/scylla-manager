// Copyright (C) 2017 ScyllaDB

package main

import (
	"github.com/scylladb/scylla-manager/pkg/command/backup/backupdelete"
	"github.com/scylladb/scylla-manager/pkg/command/backup/backupfiles"
	"github.com/scylladb/scylla-manager/pkg/command/backup/backuplist"
	"github.com/scylladb/scylla-manager/pkg/command/backup/backupvalidate"
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
	rootCmd.AddCommand(n)
}
