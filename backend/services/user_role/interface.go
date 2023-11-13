package user_role

import (
	"orb-api/models"
	"orb-api/repositories/user_role"
)

type (
	Service struct {
		UserRoleRepo *user_role.Repository
	}

	Interface interface {
		AssigneRole(IdUser uint, IdRole uint) (*models.UserRole, error)
		UnassignRole(IdUser uint, IdRole uint) (*models.UserRole, error)
	}
)
