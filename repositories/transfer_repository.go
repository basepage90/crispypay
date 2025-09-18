package repositories

import (
	"errors"
	"fmt"

	"crispypay.com/challenge/api/contract"
	"crispypay.com/challenge/infra"
	"crispypay.com/challenge/models"
	"github.com/jinzhu/gorm"
)

type TransferRepository struct {
	infra.Database
}

func NewTransferRepository(db infra.Database) TransferRepository {
	return TransferRepository{db}
}

func (r *TransferRepository) CreateTransfer(transfer *models.Transfer) error {
	return r.Create(transfer).Error
}

func (r *TransferRepository) FindByPartnerTransactionID(partnerTransactionID string) *models.Transfer {
	var transfer models.Transfer

	if err := r.Where("partner_transaction_id = ?", partnerTransactionID).First(&transfer).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			fmt.Printf("DB 조회 오류: %v\n", err)
		}
		return nil
	}

	return &transfer
}

func (r *TransferRepository) FindRecentWithUsers(limit int) ([]contract.TransferJoined, error) {
	if limit <= 0 {
		limit = 100
	}

	var rows []contract.TransferJoined

	err := r.
		Table("transfers AS t").
		Joins(`
			LEFT JOIN users AS us
			ON us.partner_transaction_id = t.partner_transaction_id
			AND us.transaction_type = ?
		`, "SENDER").
		Joins(`
			LEFT JOIN users AS ur
			ON ur.partner_transaction_id = t.partner_transaction_id
			AND ur.transaction_type = ?
		`, "RECEIVER").
		Select(`
			t.partner_transaction_id,
			t.amount, t.currency, t.status, t.fee, t.completed_at, t.created_at,
			us.name  AS sender_name,   us.phone  AS sender_phone,   us.address  AS sender_address,
			ur.name  AS receiver_name, ur.phone  AS receiver_phone, ur.address  AS receiver_address
		`).
		Order("t.created_at DESC").
		Order("t.id DESC").
		Limit(limit).
		Scan(&rows).Error

	return rows, err
}
