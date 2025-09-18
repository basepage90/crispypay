package repositories

import (
	"crispypay.com/challenge/infra"
	"crispypay.com/challenge/models"
)

type StatusHistoryRepository struct {
	infra.Database
}

func NewStatusHistoryRepository(db infra.Database) StatusHistoryRepository {
	return StatusHistoryRepository{db}
}

func (r *StatusHistoryRepository) CreateStatusHistory(statusHistory *models.StatusHistory) error {
	return r.Create(statusHistory).Error
}
