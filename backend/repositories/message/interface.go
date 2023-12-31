package message

import (
	"gorm.io/gorm"
	"orb-api/models"
)

const (
	contentMaxLen = 1024
)

type (
	Repository struct {
		GetDB func() *gorm.DB
	}

	ICreate struct {
		Sender   uint
		Receiver uint
		Content  string
	}

	IReadBySender struct {
		Sender uint
	}

	IReadByReceiver struct {
		Receiver uint
	}

	IReadChat struct {
		Sender   uint
		Receiver uint
	}

	IUpdate struct {
		ID      uint
		Content string
	}

	IDelete struct {
		ID uint
	}

	Interface interface {
		Create(ICreate) (*models.Message, error)
		ReadBySender(IReadBySender) ([]models.Message, error)
		ReadByReceiver(IReadByReceiver) ([]models.Message, error)
		ReadByChat(IReadChat) ([]models.Message, error)
		Update(IUpdate) (*models.Message, error)
		Delete(IDelete) (*models.Message, error)
	}
)
