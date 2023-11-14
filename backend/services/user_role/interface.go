package user_role

import (
	"orb-api/models"
	"orb-api/repositories/user_role"
	"orb-api/repositories/user"
	"orb-api/repositories/role"
)

type (
	Service struct {
		UserRoleRepo *userrole.Repository
		UserRepo	*user.Repository
		RoleRepo *role.Repository
	}

	Interface interface {
		AssigneRole(IdUser uint, IdRole uint) (*models.UserRole, error)
		UnassignRole(IdUser uint, IdRole uint) (*models.UserRole, error)
	}
)
