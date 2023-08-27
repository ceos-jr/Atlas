package message

import (
	"errors"
	"orb-api/models"

	"gorm.io/gorm"
)

func NewMessageRepository(connection *gorm.DB) Repository {
	return Repository{
		getDB: func() *gorm.DB {
			return connection
		},
	}
}

func (r *Repository) ValidContent(content string) bool {
	if len(content) > contentMaxLen || len(content) < 0 {
		return false
	}
	return true
}

func (r *Repository) Create(createData ICreate) (*models.Message, error) {
	var message = models.Message{
		Sender:   createData.Sender,
		Receiver: createData.Receiver,
		Content:  createData.Content,
	}

	if !r.ValidContent(createData.Content) {
		return nil, errors.New("Content too long")
	}

	result := r.getDB().Create(&message)

	if result.Error != nil {
		return nil, result.Error
	}

	return &message, nil

}

func (r *Repository) GetBySender(getBySender IGetBySender) ([]models.Message, error) {
	var messagesArray []models.Message
	var messagesMap = make(map[string]interface{})

	if getBySender.Sender == nil {
		return nil, errors.New("No field to read")
	}

	messagesMap["Sender"] = getBySender.Sender

	result := r.getDB().Where(messagesMap).Find(&messagesArray)

	if result.Error != nil {
		return nil, result.Error
	}

	return messagesArray, nil
}

func (r *Repository) GetByReceiver(getByReceiver IGetByReceiver) ([]models.Message, error) {
	var messagesArray []models.Message
	var messageMap = make(map[string]interface{})

	if getByReceiver.Receiver == nil {
		return nil, errors.New("No field to read")
	}

	messageMap["Receiver"] = getByReceiver.Receiver

	result := r.getDB().Where(messageMap).Find(&messagesArray)

	if result.Error != nil {
		return nil, result.Error
	}

	return messagesArray, nil
}

func (r *Repository) GetChat(getChat IGetChat) ([]models.Message, error) {
	var messagesArray []models.Message
	var messageMap = make(map[string]interface{})

	if getChat.Receiver == nil && getChat.Sender == nil {
		return nil, errors.New("No fields to read")
	}

	if getChat.Sender != nil {
		messageMap["Sender"] = getChat.Sender
	}

	if getChat.Receiver != nil {
		messageMap["Receiver"] = getChat.Receiver
	}

	result := r.getDB().Where(messageMap).Find(&messagesArray)

	if result.Error != nil {
		return nil, result.Error
	}

	return messagesArray, nil

}
