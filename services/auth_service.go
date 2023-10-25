package services

import (
	"errors"
	"gorm-practice/models"
	"gorm-practice/repositories"
	"gorm-practice/utils/exception"
	"gorm-practice/utils/security"

	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	CreateNewCustomer(request *models.Customer) (*models.Customer, error)
	Login(username string, password string) (string, error)
}

type authService struct {
	authRepo     repositories.AuthRepository
	customerRepo repositories.CustomerRepository
}

func NewAuthService(authRepo repositories.AuthRepository, customerRepo repositories.CustomerRepository) AuthService {
	return &authService{authRepo: authRepo, customerRepo: customerRepo}
}

func (s *authService) CreateNewCustomer(request *models.Customer) (*models.Customer, error) {

	customers, _ := s.customerRepo.List()

	for _, customer := range customers {
		if customer.UserName == request.UserName {
			return nil, exception.ErrUserNameExist
		}

		if customer.Email == request.Email {
			return nil, exception.ErrEmailExist
		}

	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), 10)

	if err != nil {
		return nil, exception.ErrFailedGeneratePassword
	}

	request.Password = string(hashedPassword)

	customer, err := s.authRepo.Create(request)

	if err != nil {
		return nil, exception.ErrFailedCreate
	}

	return customer, err
}

func (s *authService) Login(username string, password string) (string, error) {

	customer, err := s.authRepo.GetUsernamePassword(username, password)

	if err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return "", exception.ErrInvalidUsernamePassword
		}

		return "", exception.ErrCustomerDoesntExist
	}

	token, err := security.CreateAccessToken(customer)

	if err != nil {
		return "", err
	}

	return token, nil
}
