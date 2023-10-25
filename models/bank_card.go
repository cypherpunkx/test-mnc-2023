package models

import (
	"time"

	"github.com/google/uuid"
)

type BankCard struct {
	ID             uuid.UUID `gorm:"type:uuid;not null;primaryKey" json:"id"`
	UserID         string    `gorm:"type:uuid;not null" json:"userID"`
	BankName       string    `gorm:"type:varchar(255);not null" json:"bankName" binding:"required,oneof=bca mandiri bri bni"`
	CardNumber     string    `gorm:"type:varchar(255);size:6;not null;unique" json:"cardNumber" binding:"required,numeric,len=6"`
	CardholderName string    `gorm:"type:varchar(255);not null;" json:"cardholderName" binding:"required"`
	ExpirationDate time.Time `gorm:"type:timestamp;not null" json:"expirationDate" binding:"required"`
	Balance        float64   `gorm:"type:float;not null;default:0;check:Balance >= 0" json:"balance"`
	CVV            string    `gorm:"type:varchar(255);size:3" json:"cvv" binding:"required,numeric,len=3"`
}
