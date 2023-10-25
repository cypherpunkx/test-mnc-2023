package managers

import "gorm-practice/repositories"

type RepoManager interface {
	CustomerRepo() repositories.CustomerRepository
	AdminRepo() repositories.AdminRepository
	AuthRepo() repositories.AuthRepository
	TransactionRepo() repositories.TransactionRepository
}

type repoManager struct {
	infraManager InfraManager
}

func NewRepoManager(infraManager InfraManager) RepoManager {
	return &repoManager{infraManager: infraManager}
}

func (m *repoManager) CustomerRepo() repositories.CustomerRepository {
	return repositories.NewCustomerRepository(m.infraManager.Conn())
}

func (m *repoManager) AdminRepo() repositories.AdminRepository {
	return repositories.NewAdminRepository(m.infraManager.Conn())
}

func (m *repoManager) AuthRepo() repositories.AuthRepository {
	return repositories.NewAuthRepository(m.infraManager.Conn())
}

func (m *repoManager) TransactionRepo() repositories.TransactionRepository {
	return repositories.NewTransactionRepository(m.infraManager.Conn())
}
