package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Customer struct {
	ID        uuid.UUID `gorm:"type:uuid;not null;unique;primaryKey" json:"id" binding:"required"`
	UserName  string    `gorm:"type:varchar(255);not null;unique" json:"userName" binding:"required,min=5"`
	Email     string    `gorm:"type:varchar(255);not null;unique" json:"email" binding:"required,email"`
	Password  string    `gorm:"type:varchar(255);not null" json:"password" binding:"required,alphanum,min=5"`
	FirstName string    `gorm:"type:varchar(255);not null" json:"firstName" binding:"required,alpha"`
	LastName  string    `gorm:"type:varchar(255);not null" json:"lastName" binding:"required,alpha"`
	BankCard  BankCard  `gorm:"foreignKey:UserID;references:ID" json:"bankCard"`
	LastLogin time.Time `json:"lastLogin"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
