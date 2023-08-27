package message

import (
	"orb-api/models"

	"gorm.io/gorm"
)

const (
	contentMaxLen = 1024
)

type (
	Repository struct {
		getDB func() *gorm.DB
	}

	ICreate struct {
		Sender   uint
		Receiver uint
		Content  string
	}

	IGetBySender struct {
		Sender *uint
	}

	IGetByReceiver struct {
		Receiver *uint
	}

	IGetChat struct {
		Sender   *uint
		Receiver *uint
	}

	Interface interface {
		GetBySender(IGetBySender) ([]models.Message, error)
		GetByReceiver(IGetByReceiver) ([]models.Message, error)
		GetChat(IGetChat) ([]models.Message, error)
	}
)
