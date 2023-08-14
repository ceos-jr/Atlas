package user

import (
	"gorm.io/gorm"
	"orb-api/models"
	"time"
)

const (
	emailMaxLen    = 128
	emailMinLen    = 3
	nameMaxLen     = 128
	nameMinLen     = 5
	passwordMinLen = 8
)

type (
	Repository struct {
		GetDB func() *gorm.DB
	}

	ICreate struct {
		Name     string
		Email    string
		Status   uint
		Password string
	}

	IReadAll struct {
		Limit *int
	}

	IReadBy struct {
		ID        *uint
		Name      *string
		Email     *string
		Status    *uint
		Limit     *int
	}

	IUpdate struct {
		ID     uint
		Name   *string
		Email  *string
		Status *uint
	}

	IDelete struct {
		ID uint
	}

	Interface interface {
		Create(ICreate) error
		ReadAll(IReadAll) ([]models.User, error)
		ReadBy(IReadBy) ([]models.User, error)
		Update(IUpdate) error
		Delete(IDelete) error
	}
)
