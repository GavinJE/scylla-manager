// Copyright (C) 2017 ScyllaDB

package flag

import (
	"os"

	"github.com/scylladb/go-set/strset"
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
	"short",
	"long",
)

func (w Wrapper) MustSetUsages(b []byte) {
	u := make(map[string]string)
	if err := yaml.Unmarshal(b, u); err != nil {
		panic(err)
	}

	for k, v := range u {
		if keywords.Has(k) {
			continue
		}
		f := w.fs.Lookup(k)
		if f == nil {
			panic("missing flag " + k)
		}
		f.Usage = v
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
