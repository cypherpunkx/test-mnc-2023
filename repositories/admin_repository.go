package repositories

import (
	"gorm-practice/models"

	"gorm.io/gorm"
)

type AdminRepository interface {
	List() ([]*models.Customer, error)
}

type adminRepository struct {
	db *gorm.DB
}

func NewAdminRepository(db *gorm.DB) AdminRepository {
	return &adminRepository{db: db}
}

func (r *adminRepository) List() ([]*models.Customer, error) {
	customers := []*models.Customer{}

	if err := r.db.Preload("Wallet").Find(&customers).Error; err != nil {
		return nil, err
	}

	return customers, nil
}
