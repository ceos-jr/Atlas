package userrole

import (
	"gorm.io/gorm"
	"orb-api/models"
)

type (
	Repository struct {
		GetDB func() *gorm.DB
	}

	ICreate struct {
		RoleID uint
		UserID uint
	}

	IReadBy struct {
		RoleID *uint
		UserID *uint
	}

	IUpdate struct {
		UserRoleID uint
		UserID     *uint
		RoleID     *uint
	}

	IDelete struct {
		UserRoleID uint
	}

	Interface interface {
		Create(ICreate) error
		ReadAll() ([]models.UserRole, error)
		ReadBy(IReadBy) ([]models.UserRole, error)
		Update(IUpdate) error
		Delete(IDelete) error
	}
)
