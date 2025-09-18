package main

import (
	"crispypay.com/challenge/api/controllers"
	"crispypay.com/challenge/api/routes"
	"crispypay.com/challenge/infra"
	"crispypay.com/challenge/repositories"
	"crispypay.com/challenge/services"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
)

func main() {
	app := fx.New(
		infra.Module,
		routes.Module,
		controllers.Module,
		services.Module,
		repositories.Module,
		fx.WithLogger(func(logger *zap.Logger) fxevent.Logger {
			return &fxevent.ZapLogger{Logger: logger}
		}),
	)
	app.Run()
}
