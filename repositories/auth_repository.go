package repositories

import (
	"gorm-practice/models"
	"gorm-practice/utils/common"
	"gorm-practice/utils/constants"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthRepository interface {
	Create(request *models.Customer) (*models.Customer, error)
	GetUsernamePassword(username, password string) (*models.Customer, error)
}

type authRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) AuthRepository {
	return &authRepository{db: db}
}

func (r *authRepository) Create(request *models.Customer) (*models.Customer, error) {
	customer := &models.Customer{
		ID:        request.ID,
		UserName:  request.UserName,
		Email:     request.Email,
		Password:  request.Password,
		FirstName: request.FirstName,
		LastName:  request.LastName,
		LastLogin: time.Now(),
		BankCard: models.BankCard{
			ID:             common.GenerateUUID(),
			BankName:       request.BankCard.BankName,
			CardNumber:     request.BankCard.CardNumber,
			CardholderName: request.BankCard.CardholderName,
			ExpirationDate: request.BankCard.ExpirationDate,
			CVV:            request.BankCard.CVV,
		},
	}

	if err := r.db.Create(&customer).Error; err != nil {
		return nil, err
	}

	return customer, nil
}

func (r *authRepository) GetUsernamePassword(username, password string) (*models.Customer, error) {
	var customer *models.Customer

	if err := r.db.Where(constants.WHERE_USERNAME, username).Select("user_name", "password").First(&customer).Error; err != nil {
		return nil, err
	}

	err := bcrypt.CompareHashAndPassword([]byte(customer.Password), []byte(password))

	if err != nil {
		return nil, err
	}

	return customer, nil
}
