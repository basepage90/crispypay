package routes

import (
	"crispypay.com/challenge/api/controllers"
	"crispypay.com/challenge/infra"
)

type TransferRoutes struct {
	webserver  *infra.Webserver
	controller *controllers.TransferController
}

func NewTransferRoutes(
	webserver *infra.Webserver,
	controller *controllers.TransferController,

) *TransferRoutes {
	return &TransferRoutes{
		webserver:  webserver,
		controller: controller,
	}
}
func (r *TransferRoutes) Setup() {
	api := r.webserver.Gin.Group("/api")

	// 송금 요청
	api.POST("/transfer", r.controller.Transfer)

	// 단건 조회
	api.GET("/transfer/:partner_tx_id", r.controller.SearchTransfer)

	// 목록 조회
	api.GET("/transfers", r.controller.SearchTransfers)

	// 송금 취소
	api.PUT("/transfer/cancel/:partner_tx_id", r.controller.Cancel)

	// 콜백
	api.POST("/v1/webhooks/partner-callback", r.controller.Complete)
}
