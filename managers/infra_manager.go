package managers

import (
	"gorm-practice/config"

	"gorm.io/gorm"
)

type InfraManager interface {
	Conn() *gorm.DB
}

type infraManager struct {
	cfg *config.Configuration
}

func (*infraManager) Conn() *gorm.DB {
	return config.DB
}

func NewInfraManager(cfg *config.Configuration) InfraManager {
	infra := &infraManager{
		cfg: cfg,
	}

	return infra
}
