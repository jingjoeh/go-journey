package solution

import (
	"context"
	"errors"
	"net/http"
	"time"
)

type Server interface {
	ListenAndServe() error
	Shutdown(context.Context) error
}

func Run(ctx context.Context, server Server, shutdownTimeout time.Duration) error {
	errs := make(chan error, 1)
	go func() { errs <- server.ListenAndServe() }()
	select {
	case err := <-errs:
		if errors.Is(err, http.ErrServerClosed) {
			return nil
		}
		return err
	case <-ctx.Done():
		shutdownCtx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
		defer cancel()
		if err := server.Shutdown(shutdownCtx); err != nil {
			return err
		}
		err := <-errs
		if errors.Is(err, http.ErrServerClosed) {
			return nil
		}
		return err
	}
}
