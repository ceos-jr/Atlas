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

// len(content) == 0
func (r *Repository) ValidContent(content string) bool {
	if len(content) > contentMaxLen || len(content) < 0 {
		return false
	}
	return true
}

// criar uma nova condição : "content too long" e "content cannot be empty"
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

func (r *Repository) ReadBySender(readBySender IReadBySender) ([]models.Message, error) {
	var messagesArray []models.Message
	var messagesMap = make(map[string]interface{})

	if readBySender.Sender == nil {
		return nil, errors.New("No field to read")
	}

	messagesMap["Sender"] = readBySender.Sender

	result := r.getDB().Where(messagesMap).Find(&messagesArray)

	if result.Error != nil {
		return nil, result.Error
	}

	return messagesArray, nil
}

func (r *Repository) ReadByReceiver(readByReceiver IReadByReceiver) ([]models.Message, error) {
	var messagesArray []models.Message
	var messageMap = make(map[string]interface{})

	if readByReceiver.Receiver == nil {
		return nil, errors.New("No field to read")
	}

	messageMap["Receiver"] = readByReceiver.Receiver

	result := r.getDB().Where(messageMap).Find(&messagesArray)

	if result.Error != nil {
		return nil, result.Error
	}

	return messagesArray, nil
}

func (r *Repository) ReadChat(readChat IReadChat) ([]models.Message, error) {
	var messagesArray []models.Message
	var messageMap = make(map[string]interface{})

	if readChat.Receiver == nil && readChat.Sender == nil {
		return nil, errors.New("No fields to read")
	}

	if readChat.Sender != nil {
		messageMap["Sender"] = readChat.Sender
	}

	if readChat.Receiver != nil {
		messageMap["Receiver"] = readChat.Receiver
	}

	result := r.getDB().Where(messageMap).Find(&messagesArray)

	if result.Error != nil {
		return nil, result.Error
	}

	return messagesArray, nil
}

func (r *Repository) Update(updateData IUpdate) (*models.Message, error) {
	var message = models.Message{ID: updateData.ID}
	verifyExistence := r.getDB().First(&message)

	if verifyExistence.Error != nil {
		return nil, verifyExistence.Error
	}

	message.Content = *updateData.Content
	saveResult := r.getDB().Save(&message)

	if saveResult.Error != nil {
		return nil, saveResult.Error
	}

	return &message, nil
}

func (r *Repository) Delete(deleteData IDelete) (*models.Message, error) {
	var message = models.Message{ID: deleteData.ID}

	verifyExistence := r.getDB().First(&message)

	if verifyExistence.Error != nil {
		return nil, verifyExistence.Error
	}

	result := r.getDB().Delete(&message)

	if result.Error != nil {
		return nil, result.Error
	}

	return &message, nil
}
