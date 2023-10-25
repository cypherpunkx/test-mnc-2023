package services

import (
	"gorm-practice/models"
	"gorm-practice/repositories"
	"gorm-practice/utils/exception"
)

type CustomerService interface {
	GetProfile(username string) (*models.Customer, error)
	GetEmailPassword(email, password string) (*models.Customer, error)
	GetUserNamePassword(username, password string) (*models.Customer, error)
}

type customerService struct {
	customerRepo repositories.CustomerRepository
}

func NewCustomerService(customerRepo repositories.CustomerRepository) CustomerService {
	return &customerService{customerRepo: customerRepo}
}

func (s *customerService) GetUserNamePassword(username, password string) (*models.Customer, error) {
	return nil, nil
}

func (s *customerService) GetProfile(username string) (*models.Customer, error) {
	customer, err := s.customerRepo.GetByUsername(username)

	if err != nil {
		return nil, exception.ErrCustomerDoesntExist
	}

	return customer, nil
}

func (s *customerService) GetEmailPassword(email, password string) (*models.Customer, error) {
	return nil, nil
}
