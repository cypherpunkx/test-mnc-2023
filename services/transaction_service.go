package services

import (
	"gorm-practice/models"
	"gorm-practice/repositories"
	"gorm-practice/utils/exception"

	"gorm.io/gorm"
)

type TransactionService interface {
	SendMoneyToFriend(request *models.Transaction, customerID, username string) (*models.Transaction, error)
	GetHistory(username string) ([]*models.Transaction, error)
}

type transactionService struct {
	customerRepo    repositories.CustomerRepository
	transactionRepo repositories.TransactionRepository
}

func NewTransactionService(customerRepo repositories.CustomerRepository, transactionRepo repositories.TransactionRepository) TransactionService {
	return &transactionService{customerRepo: customerRepo, transactionRepo: transactionRepo}
}

func (s *transactionService) SendMoneyToFriend(request *models.Transaction, customerID, username string) (*models.Transaction, error) {

	my, err := s.customerRepo.GetByUsername(username)

	if err != nil {
		return nil, exception.ErrCustomerDoesntExist
	}

	friend, err := s.customerRepo.Get(customerID)

	if err != nil {
		return nil, exception.ErrCustomerDoesntExist
	}

	if friend.ID.String() == my.ID.String() {
		return nil, exception.ErrInvalidFriend
	}

	if request.Amount > my.BankCard.Balance {
		return nil, exception.ErrNotEnoughBalance
	}

	request.UserID = my.ID.String()

	transaction, err := s.transactionRepo.Create(request, friend.ID.String())

	if err != nil {
		return nil, gorm.ErrInvalidTransaction
	}

	return transaction, err
}

func (s *transactionService) GetHistory(username string) ([]*models.Transaction, error) {

	my, err := s.customerRepo.GetByUsername(username)

	if err != nil {
		return nil, exception.ErrCustomerDoesntExist
	}

	transactions, err := s.transactionRepo.List(my.ID.String())

	if err != nil {
		return nil, gorm.ErrRecordNotFound
	}

	return transactions, err
}
