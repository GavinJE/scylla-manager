// Copyright (C) 2017 ScyllaDB

package main

import (
	"github.com/scylladb/scylla-manager/pkg/command/backup/backupdelete"
	"github.com/scylladb/scylla-manager/pkg/command/backup/backupfiles"
	"github.com/scylladb/scylla-manager/pkg/command/backup/backuplist"
	"github.com/scylladb/scylla-manager/pkg/command/backup/backupvalidate"
)

func init() {
	rootCmd.AddCommand(backupvalidate.NewCommand(&client))
	rootCmd.AddCommand(backuplist.NewCommand(&client))
	rootCmd.AddCommand(backupfiles.NewCommand(&client))
	rootCmd.AddCommand(backupdelete.NewCommand(&client))
}
