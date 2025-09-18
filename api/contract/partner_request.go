package contract

type PartnerRequest struct {
	ExternalID  string   `json:"external_id" binding:"required"`
	Amount      string   `json:"amount" binding:"required"`
	Currency    string   `json:"currency" binding:"required,uppercase"`
	Sender      UserInfo `json:"sender" binding:"required"`
	Receiver    UserInfo `json:"receiver" binding:"required"`
	CallbackURL string   `json:"callback_url" binding:"required"`
}
