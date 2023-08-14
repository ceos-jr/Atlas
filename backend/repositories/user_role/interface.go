package userrole

import (
	"gorm.io/gorm"
	"orb-api/models"
)

type (
	Repository struct {
		GetDB func() *gorm.DB
	}

	ICreateUserRole struct {
		RoleID uint
		UserID uint
	}

	IReadBy struct {
		RoleID *uint
		UserID *uint
	}

	IUpdateUserRole struct {
		UserRoleID uint
		UserID     *uint
		RoleID     *uint
	}

	IDeleteUserRole struct {
		UserRoleID uint
	}

	Interface interface {
		Create(ICreateUserRole) error
		ReadAll() ([]models.UserRole, error)
		ReadBy(IReadBy) ([]models.UserRole, error)
		Update(IUpdateUserRole) error
		Delete(IDeleteUserRole) error
	}
)
