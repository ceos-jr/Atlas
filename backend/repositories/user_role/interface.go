package UserRolesRepo

import (
	"orb-api/models"
)

type (
	ICreateUserRole struct {
		RoleId uint
		UserId uint
	}

	IReadByUser struct {
		UserId uint
	}

	IReadByRole struct {
		RoleId uint
	}

	IUpdateUserRole struct {
		UserRoleId uint
	}

	IDeleteUserRole struct {
		UserRoleId uint
	}

	IUserRepository interface {
		Create(ICreateUserRole) error
		ReadAll() (*[]models.UserRole, error)
		ReadByRole(IReadByRole) (*[]models.UserRole, error)
		ReadByUser(IReadByUser) (*[]models.UserRole, error)
		Update(IUpdateUserRole) error
		Delete(IDeleteUserRole) error
	}
)
