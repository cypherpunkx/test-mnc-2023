package repositories

import (
	"gorm-practice/models"
	"gorm-practice/utils/constants"

	"gorm.io/gorm"
)

type TransactionRepository interface {
	Create(request *models.Transaction, friendID string) (*models.Transaction, error)
	List(id string) ([]*models.Transaction, error)
}

type transactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) TransactionRepository {
	return &transactionRepository{db: db}
}

func (r *transactionRepository) Create(request *models.Transaction, friendID string) (*models.Transaction, error) {
	transaction := models.Transaction{
		ID:              request.ID,
		UserID:          request.UserID,
		TransactionType: request.TransactionType,
		Amount:          request.Amount,
		Description:     request.Description,
		Timestamp:       request.Timestamp,
	}

	r.db.Transaction(func(tx *gorm.DB) error {
		var bankCard models.BankCard

		if transaction.TransactionType == "send" {
			if err := tx.Where(constants.WHERE_USER_ID, transaction.UserID).Select("balance").First(&bankCard).Error; err != nil {
				return gorm.ErrInvalidTransaction
			}

			bankCard.Balance -= transaction.Amount

			if err := tx.Where(constants.WHERE_USER_ID, transaction.UserID).Select("balance").Updates(&bankCard).Error; err != nil {
				return gorm.ErrInvalidTransaction
			}

			if err := tx.Where(constants.WHERE_USER_ID, friendID).Select("balance").First(&bankCard).Error; err != nil {
				return gorm.ErrInvalidTransaction
			}

			bankCard.Balance += transaction.Amount

			if err := tx.Where(constants.WHERE_USER_ID, friendID).Select("balance").Updates(&bankCard).Error; err != nil {
				return gorm.ErrInvalidTransaction
			}

			if err := tx.Create(&transaction).Error; err != nil {
				return gorm.ErrInvalidTransaction
			}
		} else {
			return gorm.ErrInvalidTransaction
		}

		return nil
	})

	return &transaction, nil
}

func (r *transactionRepository) List(id string) ([]*models.Transaction, error) {
	var transactions []*models.Transaction

	if err := r.db.Where(constants.WHERE_USER_ID, id).Find(&transactions).Error; err != nil {
		return nil, err
	}

	return transactions, nil
}
