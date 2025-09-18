package contract

import "time"

type TransferJoined struct {
	PartnerTransactionID string
	Amount               string
	Currency             string
	Status               string
	Fee                  string
	CompletedAt          string
	CreatedAt            time.Time

	SenderName    string
	SenderPhone   string
	SenderAddress string

	ReceiverName    string
	ReceiverPhone   string
	ReceiverAddress string
}
