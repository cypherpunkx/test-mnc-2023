package managers

import (
	"gorm-practice/services"
)

type ServiceManager interface {
	CustomerService() services.CustomerService
	AdminService() services.AdminService
	AuthService() services.AuthService
	TransactionService() services.TransactionService
}

type serviceManager struct {
	repoManager RepoManager
}

func NewServiceManager(repoManager RepoManager) ServiceManager {
	return &serviceManager{repoManager: repoManager}
}

func (m *serviceManager) CustomerService() services.CustomerService {
	return services.NewCustomerService(m.repoManager.CustomerRepo())
}

func (m *serviceManager) AdminService() services.AdminService {
	return services.NewAdminService(m.repoManager.AdminRepo())
}

func (m *serviceManager) AuthService() services.AuthService {
	return services.NewAuthService(m.repoManager.AuthRepo(), m.repoManager.CustomerRepo())
}

func (m *serviceManager) TransactionService() services.TransactionService {
	return services.NewTransactionService(m.repoManager.CustomerRepo(), m.repoManager.TransactionRepo())
}
