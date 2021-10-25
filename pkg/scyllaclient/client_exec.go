// Copyright (C) 2017 ScyllaDB

package scyllaclient

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"sync"
	"time"

	"github.com/goware/prefixer"
	"github.com/pkg/errors"
	"github.com/scylladb/scylla-manager/pkg/util/parallel"
)

func (c *Client) Exec(ctx context.Context, host []string, limit int, stdin []byte, stdout io.Writer) error {
	ctx = customTimeout(ctx, time.Hour)

	r := make([]io.Reader, len(host))
	for i := range host {
		r[i] = bytes.NewReader(stdin)
	}

	stdout = &syncStdout{w: stdout}
	return parallel.Run(len(host), limit, func(i int) error {
		return c.ExecHost(ctx, host[i], host[i]+"| ", r[i], stdout)
	})
}

func (c *Client) ExecHost(ctx context.Context, host, prefix string, stdin io.Reader, stdout io.Writer) error {
	// Due to OpenAPI limitations we manually construct and sent the request
	// object to stream process the response body.
	const urlPath = "/exec"

	u := c.newURL(host, urlPath)
	req, err := http.NewRequestWithContext(forceHost(ctx, host), http.MethodPost, u.String(), stdin)
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "application/octet-stream")

	resp, err := c.transport.RoundTrip(req)
	if err != nil {
		return errors.Wrap(err, "round trip")
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return makeAgentError(resp)
	}

	_, err = io.Copy(stdout, prefixer.New(resp.Body, prefix))
	if errors.Is(err, io.EOF) {
		err = nil
	}
	return err
}

type syncStdout struct {
	mu sync.Mutex
	w  io.Writer
}

func (s *syncStdout) Write(p []byte) (n int, err error) {
	s.mu.Lock()
	n, err = s.w.Write(p)
	s.mu.Unlock()
	return
}