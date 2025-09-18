package repositories

import (
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(NewTransferRepository),
	fx.Provide(NewUserRepository),
	fx.Provide(NewStatusHistoryRepository),
)
