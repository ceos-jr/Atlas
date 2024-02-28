package sector

import (
	"orb-api/models"

	"gorm.io/gorm"
)

const (
	nameMaxLen     = 128
	nameMinLen     = 3
)

type (
	Repository struct {
		GetDB func() *gorm.DB
	}

	ICreate struct {
		Name        string
		Description string
		AdmID       uint
	}

	IReadAll struct {
		Limit *int
	}

	IReadBy struct {
		ID    *uint
		Name  *string
		AdmID *uint
		Limit *int
	}

	IUpdate struct {
		ID          uint
		Name        *string
		Description *string
		AdmID       *uint
		Members     *[]uint
		Projects    *[]uint
	}

	IDelete struct {
		ID uint
	}

	Interface interface {
		Create(ICreate) (*models.Sector, error)
		ReadAll() ([]models.Sector, error)
		ReadBy(IReadBy) ([]models.Sector, error)
		Update(IUpdate) (*models.Sector, error)
		Delete(IDelete) (*models.Sector, error)
	}
)
