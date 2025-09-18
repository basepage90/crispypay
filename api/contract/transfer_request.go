package contract

type TransferRequest struct {
	Sender   UserInfo `json:"sender" binding:"required"`
	Receiver UserInfo `json:"receiver" binding:"required"`

	Amount             string `json:"amount" binding:"required"`
	Currency           string `json:"currency" binding:"required,uppercase"`
	DestinationCountry string `json:"destination_country" binding:"required,uppercase"`
	PartnerProvider    string `json:"partner_provider" binding:"required"`
	Purpose            string `json:"purpose" binding:"required"`
}
