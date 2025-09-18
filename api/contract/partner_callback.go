package contract

type PartnerCallback struct {
	PartnerTransactionID string `json:"partner_transaction_id" binding:"required"`
	RemittanceID         string `json:"remittance_id" binding:"required"`
	Status               string `json:"status" binding:"required"`
	CompletedAt          string `json:"completed_at" binding:"required"`
	FailureReason        string `json:"failure_reason" binding:"required"`
	PartnerProvider      string `json:"partner_provider" binding:"required"`
}
