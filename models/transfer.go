package models

import "time"

type Transfer struct {
	ID                 uint   `gorm:"primary_key;column:id"`
	ExternalID         string `gorm:"column:external_id"`
	Amount             string `gorm:"column:amount"`
	Currency           string `gorm:"column:currency"`
	DestinationCountry string `gorm:"column:destination_country"`
	PartnerProvider    string `gorm:"column:partner_provider"`
	Purpose            string `gorm:"column:purpose"`
	CallbackURL        string `gorm:"column:callback_url"`

	PartnerTransactionID string `gorm:"column:partner_transaction_id"`
	Status               string `gorm:"column:status"`
	EstimatedCompletion  string `gorm:"column:estimated_completion"`
	CompletedAt          string `gorm:"column:completed_at"`
	Fee                  string `gorm:"column:fee"`
	Message              string `gorm:"column:message"`

	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}
