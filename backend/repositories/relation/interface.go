package relation

import (
	"gorm.io/gorm"
	"orb-api/models"
)

const (
	_invalidLeftKey    = "invalid LeftUserRoleId value"
	_invalidRightKey   = "invalid RightUserRoleId value"
	_invalidStrongSide = "invalid StrongSide value"
	_invalidID         = "invalid id value"

	CErrorInvalidLeftKey    = _invalidLeftKey
	CErrorInvalidRightKey   = _invalidRightKey
	CErrorInvalidStrongSide = _invalidStrongSide

	RErrorEmptyReadBy       = "no fields to search for"
	RErrorInvalidID         = _invalidID
	RErrorInvalidLeftKey    = _invalidLeftKey
	RErrorInvalidRightKey   = _invalidRightKey
	RErrorInvalidStrongSide = _invalidStrongSide

	UErrorEmptyUpdate       = "no fields to update"
	UErrorInvalidID         = _invalidID
	UErrorInvalidLeftKey    = _invalidLeftKey
	UErrorInvalidRightKey   = _invalidRightKey
	UErrorInvalidStrongSide = _invalidStrongSide

	DErrorInvalidID = _invalidID
)

type (
	Repository struct {
		GetDB func() *gorm.DB
	}

	ICreate struct {
		StrongSide      string
		RightUserRoleId uint
		LeftUserRoleId  uint
	}

	IUpdate struct {
		ID              uint
		StrongSide      *string
		RightUserRoleId *uint
		LeftUserRoleId  *uint
	}

	// implement full pagination latter: https://dev.to/rafaelgfirmino/pagination-using-gorm-scopes-3k5f
	IReadAll struct {
		Limit *uint
	}

	IReadBy struct {
		ID              *uint
		StrongSide      *string
		RightUserRoleId *uint
		LeftUserRoleId  *uint
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
