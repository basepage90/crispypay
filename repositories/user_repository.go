package repositories

import (
	"errors"
	"fmt"

	"crispypay.com/challenge/infra"
	"crispypay.com/challenge/models"
	"github.com/jinzhu/gorm"
)

type UserRepository struct {
	infra.Database
}

func NewUserRepository(db infra.Database) UserRepository {
	return UserRepository{db}
}

func (r *UserRepository) CreateUser(user *models.User) error {
	return r.Create(user).Error
}

func (r *UserRepository) FindUser(partnerTransactionID string, transactionType string) *models.User {
	var user models.User

	if err := r.Where("partner_transaction_id = ?", partnerTransactionID).First(&user).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			fmt.Printf("DB 조회 오류: %v\n", err)
		}
		return nil
	}

	return &user
}
