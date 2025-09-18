package contract

import "time"

type SearchTransferResponse struct {
	PartnerTransactionID string    `json:"partner_transaction_id" binding:"required"`
	Amount               string    `json:"amount" binding:"required"`
	Currency             string    `json:"currency" binding:"required,uppercase"`
	Sender               UserInfo  `json:"sender" binding:"required"`
	Receiver             UserInfo  `json:"receiver" binding:"required"`
	Status               string    `json:"status" binding:"required"`
	Fee                  string    `json:"fee" binding:"required"`
	CompletedAt          string    `json:"completed_at" binding:"required"`
	CreatedAt            time.Time `json:"created_at" `
}
