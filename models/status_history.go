package models

import "time"

type StatusHistory struct {
	TransferID uint      `gorm:"column:transfer_id"`
	Status     string    `gorm:"column:status"`
	CreatedAt  time.Time `gorm:"column:created_at"`
}
