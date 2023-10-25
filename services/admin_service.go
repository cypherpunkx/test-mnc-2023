package services

import (
	"gorm-practice/models"
	"gorm-practice/repositories"

	"gorm.io/gorm"
)

type AdminService interface {
	GetAllUsers() ([]*models.Customer, error)
}

type adminService struct {
	adminRepo repositories.AdminRepository
}

func NewAdminService(adminRepo repositories.AdminRepository) AdminService {
	return &adminService{adminRepo: adminRepo}
}

func (s *adminService) GetAllUsers() ([]*models.Customer, error) {

	users, err := s.adminRepo.List()

	if err != nil {
		return nil, gorm.ErrRecordNotFound
	}

	return users, err
}
