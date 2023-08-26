package role

import (
	"gorm.io/gorm"
	"orb-api/models"
)

type (
	Repository struct {
		getDB func() *gorm.DB
	}

	ICreate struct {
		Name        string
		Description string
	}

	IReadBy struct {
		ID          *uint
		Name        *string
		Description *string
	}

	IReadAll struct {
		Limit *int
	}

	IUpdate struct {
		RoleID      uint
		Name        *string
		Description *string
	}

	IDelete struct {
		RoleID uint
	}

	Interface interface {
		Create(ICreate) (*models.Role, error)
		ReadAll() ([]models.Role, error)
		ReadBy(IReadBy) ([]models.Role, error)
		Update(IUpdate) (*models.Role, error)
		Delete(IDelete) (*models.Role, error)
	}
)
