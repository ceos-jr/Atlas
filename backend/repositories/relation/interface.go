package relation

import (
	"gorm.io/gorm"
	"orb-api/models"
)

const (
	CErrorInvalidStrongSide = "invalid StrongSide value"
	RErrorEmptyReadBy       = "no fields to search for"
	RErrorInvalidStrongSide = "invalid StrongSide value"
	UErrorEmptyUpdate       = "no fields to update"
	UErrorInvalidStrongSide = "invalid StrongSide value"
)

type (
	Repository struct {
		GetDB func() *gorm.DB
	}

	ICreate struct {
		StrongSide      string
		RightUserRoleID uint
		LeftUserRoleID  uint
	}

	IUpdate struct {
		ID              uint
		StrongSide      *string
		RightUserRoleID *uint
		LeftUserRoleID  *uint
	}

	IReadAll struct {
		Limit *uint
	}

	IReadBy struct {
		ID              *uint
		StrongSide      *string
		RightUserRoleID *uint
		LeftUserRoleID  *uint
		Limit           *uint
	}

	IDelete struct {
		ID uint
	}

	Interface interface {
		Create(ICreate) (*models.Relation, error)
		ReadAll(IReadAll) ([]models.Relation, error)
		ReadBy(IReadBy) ([]models.Relation, error)
		Update(IUpdate) (*models.Relation, error)
		Delete(IDelete) (*models.Relation, error)
	}
)
