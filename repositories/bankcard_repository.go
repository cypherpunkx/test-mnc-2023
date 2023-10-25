package repositories

import (
	"gorm-practice/models"
	"gorm-practice/utils/constants"

	"gorm.io/gorm"
)

type CardRepository interface {
	Get(id string) (*models.BankCard, error)
	GetByUserID(bankCardID, customerID string) (*models.BankCard, error)
}

type cardRepository struct {
	db *gorm.DB
}

func NewCardRepository(db *gorm.DB) CardRepository {
	return &cardRepository{db: db}
}

func (r *cardRepository) Get(id string) (*models.BankCard, error) {
	var card *models.BankCard

	if err := r.db.Where(constants.WHERE_ID, id).First(&card).Error; err != nil {
		return nil, err
	}

	return card, nil
}

func (r *cardRepository) GetByUserID(bankCardID, customerID string) (*models.BankCard, error) {
	var card *models.BankCard

	if err := r.db.Where(constants.WHERE_ID_AND_USER_ID, bankCardID, customerID).First(&card).Error; err != nil {
		return nil, err
	}

	return card, nil
}
