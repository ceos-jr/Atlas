package role

import (
	"orb-api/models"

	"gorm.io/gorm"
)

const (
	nameMaxLen        = 128
	nameMinLen        = 3
	descriptionMaxLen = 128
	descriptionMinLen = 3
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
