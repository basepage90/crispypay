package models

type User struct {
	PartnerTransactionID string `gorm:"column:partner_transaction_id"`
	TransactionType      string `gorm:"column:transaction_type"`
	Name                 string `gorm:"column:name"`
	Phone                string `gorm:"column:phone"`
	Address              string `gorm:"column:address"`
}
