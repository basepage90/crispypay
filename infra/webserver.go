package infra

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type Webserver struct {
	Gin *gin.Engine
}

func NewWebserver() *Webserver {
	g := gin.New()
	g.SetTrustedProxies(nil)
	g.Use(gin.Recovery())

	g.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "centbe transfer server is running"})
	})

	var webserver = new(Webserver)
	webserver.Gin = g

	return webserver
}

func RunHTTPServer(lc fx.Lifecycle, ws *Webserver, logger *zap.Logger) {
	srv := &http.Server{
		Addr:    ":8080",
		Handler: ws.Gin,
	}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			logger.Info("starting HTTP server", zap.String("addr", srv.Addr))
			go func() {
				if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
					logger.Fatal("ListenAndServe failed", zap.Error(err))
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			logger.Info("stopping HTTP server")
			shutdownCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
			defer cancel()
			return srv.Shutdown(shutdownCtx)
		},
	})
}
