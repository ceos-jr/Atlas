package task

import (
	"gorm.io/gorm"
	"orb-api/models"
	"time"
)

type (
	Repository struct {
		GetDB func() *gorm.DB
	}

	ICreate struct {
		Description string
		AssignedTo  uint
		CreatedBy   uint
		Status      uint
		Deadline    time.Time
	}

	IReadBy struct {
		ID         *uint
		AssignedTo *uint
		CreatedBy  *uint
		Status     *uint
		TimeRange  *time.Time
	}

	IUpdate struct {
		ID          uint
		Description *string
		AssignedTo  *uint
		CreatedBy   *uint
		Status      *uint
		Deadline    *time.Time
	}

	IDelete struct {
		ID uint
	}

	Interface interface {
		Create(ICreate) (*models.Task, error)
		ReadAll() ([]models.Task, error)
		ReadBy(IReadBy) ([]models.Task, error)
		Update(IUpdate) (*models.Task, error)
		Delete(IDelete) (*models.Task, error)
	}
)
