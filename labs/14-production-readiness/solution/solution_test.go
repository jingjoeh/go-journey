package solution_test

import (
	target "bootcamp/14-production-readiness/solution"
	"context"
	"net/http"
	"sync/atomic"
	"testing"
	"time"
)

type fakeServer struct {
	stopped  chan struct{}
	shutdown atomic.Bool
}

func (f *fakeServer) ListenAndServe() error { <-f.stopped; return http.ErrServerClosed }
func (f *fakeServer) Shutdown(context.Context) error {
	f.shutdown.Store(true)
	close(f.stopped)
	return nil
}
func TestRunShutdown(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	server := &fakeServer{stopped: make(chan struct{})}
	if err := target.Run(ctx, server, time.Second); err != nil {
		t.Fatal(err)
	}
	if !server.shutdown.Load() {
		t.Fatal("shutdown not called")
	}
}
