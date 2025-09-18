package contract

type PartnerResponse struct {
	PartnerTransactionID string `json:"partner_transaction_id" binding:"required"`
	Status               string `json:"status" binding:"required"`
	EstimatedCompletion  string `json:"estimated_completion" `
	Fee                  string `json:"fee" binding:"required"`
	Message              string `json:"message" binding:"required"`
}
