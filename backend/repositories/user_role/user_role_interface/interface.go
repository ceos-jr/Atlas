package userRoleInterface

import (
	"orb-api/models"
)

type (
	ICreateUserRole struct {
		RoleId uint
		UserId uint
	}

	IReadBy struct {
		RoleId *uint
		UserId *uint
	}

	IUpdateUserRole struct {
		UserRoleId uint
		UserId     *uint
		RoleId     *uint
	}

	IDeleteUserRole struct {
		UserRoleId uint
	}

	InterfaceUserRole interface {
		Create(ICreateUserRole) error
		ReadAll() (*[]models.UserRole, error)
		ReadBy(by IReadBy) (*[]models.UserRole, error)
		Update(IUpdateUserRole) error
		Delete(IDeleteUserRole) error
	}
)
