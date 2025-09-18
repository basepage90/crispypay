package controllers

import (
	"log"
	"net/http"

	"crispypay.com/challenge/api/contract"
	"crispypay.com/challenge/services"
	"github.com/gin-gonic/gin"
)

type TransferController struct {
	service *services.TransferService
}

func NewTransferController(transferService *services.TransferService) *TransferController {
	return &TransferController{
		service: transferService,
	}
}

func (c *TransferController) Transfer(ctx *gin.Context) {
	var request contract.TransferRequest

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := c.service.Transfer(request); err != nil {
		log.Printf("송금 이력 기록 실패: %v", err)
	}
}

func (c *TransferController) SearchTransfer(ctx *gin.Context) {
	partnerTxID := ctx.Param("partner_tx_id")
	if partnerTxID == "" {
		ctx.JSON(400, gin.H{"error": "partner_tx_id is required"})
		return
	}

	response := c.service.SearchTransfer(partnerTxID)

	ctx.JSON(http.StatusOK, response)
}

func (c *TransferController) SearchTransfers(ctx *gin.Context) {
	response := c.service.SearchTransfers()

	ctx.JSON(http.StatusOK, response)
}

func (c *TransferController) Cancel(ctx *gin.Context) {
	partnerTxID := ctx.Param("partner_tx_id")
	if partnerTxID == "" {
		ctx.JSON(400, gin.H{"error": "partner_tx_id is required"})
		return
	}

	if err := c.service.Cancel(partnerTxID); err != nil {
		log.Printf("취소 처리 실패: %v", err)
	}
}

func (c *TransferController) Complete(ctx *gin.Context) {
	var callback contract.PartnerCallback

	if err := ctx.ShouldBindJSON(&callback); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := c.service.CompleteTransfer(callback); err != nil {
		log.Printf("송금 이력 기록 실패: %v", err)
	}
}
