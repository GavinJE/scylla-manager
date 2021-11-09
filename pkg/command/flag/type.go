// Copyright (C) 2017 ScyllaDB

package flag

import (
	"strings"
	"time"

	"github.com/scylladb/scylla-manager/pkg/util/duration"
	"github.com/scylladb/scylla-manager/pkg/util/timeutc"
	flag "github.com/spf13/pflag"
)

type Time struct {
	time.Time
}

var _ flag.Value = (*Time)(nil)

func (t *Time) Set(s string) (err error) {
	t.Time, err = parserTime(s)
	return
}

func parserTime(s string) (time.Time, error) {
	if strings.HasPrefix(s, "now") {
		now := timeutc.Now()
		if s == "now" {
			return now, nil
		}
		d, err := duration.ParseDuration(s[3:])
		if err != nil {
			return time.Time{}, err
		}
		return now.Add(d.Duration()), nil
	}

	t, err := timeutc.Parse(time.RFC3339, s)
	if err != nil {
		return time.Time{}, err
	}
	return t.UTC(), nil
}

func (t *Time) Type() string {
	return "string"
}

type Duration struct {
	duration.Duration
}

var _ flag.Value = (*Duration)(nil)

func (d Duration) Set(s string) (err error) {
	d.Duration, err = duration.ParseDuration(s)
	return
}

func (d Duration) Type() string {
	return "string"
}
