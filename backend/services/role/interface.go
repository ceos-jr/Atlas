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
		NewRole(name, description string) (*models.Role, error)
		UpdateName(id uint, name string) (*models.Role, error)
		UpdateDescription(id uint, description string) (*models.Role, error)
	}
)
