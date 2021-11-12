// Copyright (C) 2017 ScyllaDB

package flag

import (
	"os"

	"github.com/scylladb/go-set/strset"
	"github.com/spf13/cobra"
	flag "github.com/spf13/pflag"
	"gopkg.in/yaml.v2"
)

type Wrapper struct {
	fs *flag.FlagSet
}

func Wrap(fs *flag.FlagSet) Wrapper {
	return Wrapper{fs: fs}
}

func (w Wrapper) Unwrap() *flag.FlagSet {
	return w.fs
}

var keywords = strset.New(
	"use",
	"deprecated",
	"short",
	"long",
)

func MustSetUsages(cmd *cobra.Command, b []byte, required ...string) {
	u := make(map[string]string)
	if err := yaml.Unmarshal(b, u); err != nil {
		panic(err)
	}

	fs := cmd.Flags()
	// Set usages from file
	for k, v := range u {
		if keywords.Has(k) {
			continue
		}
		f := fs.Lookup(k)
		if f == nil {
			panic("missing flag " + k)
		}
		f.Usage = v
	}
	// Make sure flags are set
	fs.Visit(func(f *flag.Flag) {
		if len(f.Usage) == 0 {
			panic("no usage for flag " + f.Name)
		}
	})
	// Mark flags as required
	for _, name := range required {
		if err := cmd.MarkFlagRequired(name); err != nil {
			panic(err)
		}
	}
}

//
// Task schedule flags
//

func (w Wrapper) Interval(p *Duration) {
	w.fs.VarP(p, "interval", "i", usage["interval"])
}

func (w Wrapper) StartDate(p *Time) {
	w.fs.VarP(p, "start-date", "s", usage["start-date"])
}

func (w Wrapper) NumRetries(p *int, def int) {
	w.fs.IntVarP(p, "num-retries", "r", def, usage["num-retries"])
}

//
// Common flags
//

func (w Wrapper) Cluster(p *string) {
	w.fs.StringVarP(p, "cluster", "c", os.Getenv("SCYLLA_MANAGER_CLUSTER"), usage["cluster"])
}

func (w Wrapper) Datacenter(p *[]string) {
	w.fs.StringSliceVar(p, "dc", nil, usage["dc"])
}

func (w Wrapper) Keyspace(p *[]string) {
	w.fs.StringSliceVarP(p, "keyspace", "K", nil, usage["keyspace"])
}

func (w Wrapper) Location(p *[]string) {
	w.fs.StringSliceVarP(p, "location", "L", nil, usage["location"])
}
