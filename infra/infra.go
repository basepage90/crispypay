package infra

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(NewWebserver),
	fx.Provide(NewLogger),
	fx.Provide(NewDatabase),
	fx.Provide(NewHTTPClient),
	fx.Invoke(RunHTTPServer),
	fx.Invoke(GetHost),
)
