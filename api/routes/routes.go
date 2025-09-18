package routes

import (
	"go.uber.org/fx"
)

type Routes []Route

type Route interface {
	Setup()
}

func NewRoutes(transferRoutes *TransferRoutes) Routes {
	return Routes{
		transferRoutes,
	}
}

func (r Routes) Setup() {
	for _, route := range r {
		route.Setup()
	}
}

var Module = fx.Options(
	fx.Provide(NewTransferRoutes),
	fx.Provide(NewRoutes),
	fx.Invoke(func(r Routes) {
		r.Setup()
	}),
)
