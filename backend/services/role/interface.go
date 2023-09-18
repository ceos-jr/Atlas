package role

import (
	"orb-api/models"
	"orb-api/repositories/role"
)

type (
	Service struct {
		RoleRepo *role.Repository
	}

	Interface interface {
		CreateNewRole(name, description string) (*models.Role, error)
		UpdateName(id uint, name string) (*models.Role, error)
		UpdateDescription(id uint, description string) (*models.Role, error)
	}
)
