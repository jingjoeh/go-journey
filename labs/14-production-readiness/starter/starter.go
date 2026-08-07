package starter

import (
	"context"
	"net/http"
	"time"
)

type Server interface {
	ListenAndServe() error
	Shutdown(context.Context) error
}

func Run(ctx context.Context, server Server, shutdownTimeout time.Duration) error {
	panic("TODO: implement Run")
}

var _ = http.ErrServerClosed
