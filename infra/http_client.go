package infra

import (
	"context"
	"net/http"
	"time"

	"go.uber.org/fx"
)

func NewHTTPClient(lc fx.Lifecycle) *http.Client {
	transport := &http.Transport{
		MaxIdleConns:        100,
		MaxIdleConnsPerHost: 10,
	}

	client := &http.Client{
		Timeout:   10 * time.Second,
		Transport: transport,
	}

	lc.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			transport.CloseIdleConnections()
			return nil
		},
	})

	return client

}
