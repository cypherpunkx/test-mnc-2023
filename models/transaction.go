package models

import (
	"time"

	"github.com/google/uuid"
)

type Transaction struct {
	ID              uuid.UUID `gorm:"type:uuid;primaryKey;not null;unique" json:"id" binding:"required"`
	UserID          string    `gorm:"type:uuid;not null" json:"UserID"`
	TransactionType string    `gorm:"varchar(255);not null" json:"transactionType" binding:"required,alpha,oneof=deposit withdrawal send payment"`
	Amount          float64   `gorm:"type:float;not null;default:0;check:Amount >= 0" json:"amount" binding:"required,numeric"`
	Description     string    `gorm:"type:text" json:"description" binding:"omitempty"`
	Timestamp       time.Time `gorm:"type:timestamp;default:current_timestamp" json:"timestamp"`
}
