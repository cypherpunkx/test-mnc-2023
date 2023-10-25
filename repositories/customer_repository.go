package repositories

import (
	"gorm-practice/models"
	"gorm-practice/utils/constants"

	"gorm.io/gorm"
)

type CustomerRepository interface {
	List() ([]*models.Customer, error)
	Get(id string) (*models.Customer, error)
	GetByUsername(username string) (*models.Customer, error)
}

type customerRepository struct {
	db *gorm.DB
}

func NewCustomerRepository(db *gorm.DB) CustomerRepository {
	return &customerRepository{db: db}
}

func (r *customerRepository) List() ([]*models.Customer, error) {
	customers := []*models.Customer{}

	if err := r.db.Find(&customers).Error; err != nil {
		return nil, err
	}

	return customers, nil
}

func (r *customerRepository) Get(id string) (*models.Customer, error) {
	customer := &models.Customer{}

	if err := r.db.Where(constants.WHERE_ID, id).First(&customer).Error; err != nil {
		return nil, err
	}

	return customer, nil
}

func (r *customerRepository) GetByUsername(username string) (*models.Customer, error) {
	customer := &models.Customer{}

	if err := r.db.Where(constants.WHERE_USERNAME, username).Preload("BankCard").First(&customer).Error; err != nil {
		return nil, err
	}

	return customer, nil
}
