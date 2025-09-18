package infra

import (
	"fmt"
	"log"

	"crispypay.com/challenge/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Database struct {
	*gorm.DB
}

func NewDatabase() Database {
	url := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", "user01", "pass123", "crispyblog.kr", "3306", "crispypay")

	db, err := gorm.Open(mysql.Open(url))
	if err != nil {
		log.Fatal("DB 연결 실패:", err)
	}

	// 서버 구동시, 테이블 드랍
	if err := db.Migrator().DropTable(&models.Transfer{}, &models.StatusHistory{}, &models.User{}); err != nil {
		log.Fatal(err)
	}

	// 서버 구동시, 테이블 생성
	if err := db.AutoMigrate(&models.Transfer{}, &models.StatusHistory{}, &models.User{}); err != nil {
		log.Fatal(err)
	}

	return Database{DB: db}
}
