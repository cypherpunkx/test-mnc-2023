package config

import (
	"gorm-practice/models"
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	var once sync.Once

	dsn := "host=localhost user=postgres password=admin dbname=db_zenith port=5432 sslmode=disable TimeZone=Asia/Shanghai"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		PrepareStmt:    true,
		TranslateError: true,
	})

	if err != nil {
		panic(err)
	}

	once.Do(func() {
		DB = db
	})

}

func SyncDB() {
	if err := DB.AutoMigrate(&models.Customer{}, &models.BankCard{}, &models.Transaction{}); err != nil {
		panic(err)
	}
}
