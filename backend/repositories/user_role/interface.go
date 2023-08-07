package userRoleRepository

import (
	"orb-api/models"
)

type (
	IReadResultUserRole struct {
		data   *[]models.UserRole
		status error
	}

	ICreateUserRole struct {
		roleId uint
		userId uint
	}

	IReadByUser struct {
		userId uint
	}

	IReadByRole struct {
		roleId uint
	}

	IUpdateUserRole struct {
		userRoleId uint
	}

	IDeleteUserRole struct {
		userRoleId uint
	}

	IUserRepository interface {
		ReadAll() IReadResultUserRole
		ReadByRole(IReadByRole) IReadResultUserRole
		ReadByUser(IReadByUser) IReadResultUserRole
		Update(IUpdateUserRole) error
		Create(ICreateUserRole) error
		Delete(IDeleteUserRole) error
	}
)
