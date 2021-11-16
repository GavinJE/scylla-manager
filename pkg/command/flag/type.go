// Copyright (C) 2017 ScyllaDB

package flag

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/pkg/errors"
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

func (t *Time) String() string {
	if t.IsZero() {
		return ""
	}
	return t.Time.String()
}

type Duration struct {
	duration.Duration
}

var _ flag.Value = (*Duration)(nil)

func (d *Duration) Set(s string) (err error) {
	d.Duration, err = duration.ParseDuration(s)
	return
}

func (d *Duration) Type() string {
	return "string"
}

func (d *Duration) String() string {
	if d.Duration == 0 {
		return ""
	}
	return d.Duration.String()
}

// Intensity represents intensity flag which is a float64 value with a custom validation.
type Intensity struct {
	Value float64
}

// String returns intensity value as string.
func (fl *Intensity) String() string {
	return fmt.Sprint(fl.Value)
}

// Set validates and sets intensity value.
func (fl *Intensity) Set(s string) error {
	errValidation := errors.New("intensity must be an integer >= 1 or a decimal between (0,1)")

	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return errValidation
	}
	if f > 1 {
		_, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			return errValidation
		}
	}

	fl.Value = f
	return nil
}

// Type returns type of intensity.
func (fl *Intensity) Type() string {
	return "float64"
}
