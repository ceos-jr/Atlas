package userRole

import (
	"gorm.io/gorm"
	"orb-api/models"
	"orb-api/repositories"
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
		newUserId  *uint
		newRoleId  *uint
	}

	IDeleteUserRole struct {
		UserRoleId uint
	}

	RUserRole struct {
		repo *repositories.Repository
	}

	RIUserRole interface {
		Setup(*gorm.DB) RUserRole
		Create(ICreateUserRole) error
		ReadAll() (*[]models.UserRole, error)
		ReadBy(by IReadBy) (*[]models.UserRole, error)
		Update(IUpdateUserRole) error
		Delete(IDeleteUserRole) error
	}
)
