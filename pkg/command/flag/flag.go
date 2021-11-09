// Copyright (C) 2017 ScyllaDB

package flag

import flag "github.com/spf13/pflag"

type Wrapper struct {
	*flag.FlagSet
}

func Wrap(fs *flag.FlagSet) Wrapper {
	return Wrapper{FlagSet: fs}
}

//
// Common task flags
//

func (w Wrapper) Interval(p *Duration) {
	w.VarP(p, "interval", "i", usage["interval"])
}

func (w Wrapper) StartDate(p *Time) {
	w.VarP(p, "start-date", "s", usage["start-date"])
}

func (w Wrapper) NumRetries(p *int64, def int64) {
	w.Int64VarP(p, "num-retries", "r", def, usage["num-retries"])
}

//
// Common flags
//

func (w Wrapper) Datacenter(p *[]string) {
	w.StringSliceVar(p, "dc", nil, usage["dc"])
}

func (w Wrapper) Keyspace(p *[]string) {
	w.StringSliceVarP(p, "keyspace", "K", nil, usage["keyspace"])
}

func (w Wrapper) Location(p *[]string) {
	w.StringSliceVarP(p, "location", "L", nil, usage["location"])
}
