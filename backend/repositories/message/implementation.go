package message

import (
	"errors"
	"orb-api/models"
	"time"

	"gorm.io/gorm"
)

func NewMessageRepository(connection *gorm.DB) Repository {
	return Repository{
		GetDB: func() *gorm.DB {
			return connection
		},
	}
}

// len(content) == 0
func (r *Repository) ValidContent(content string) bool {
	if len(content) > contentMaxLen || len(content) <= 0 {
		return false
	}
	if len(content) > contentMaxLen {
		return false
	}
	return true
}

func (r *Repository) ValidUser(id uint) bool {
	user := models.User{ID: id}

	verifyUser := r.GetDB().First(&user).Error

	if verifyUser != nil {
		return false
	}

	return true
}

func (r *Repository) Create(createData ICreate) (*models.Message, error) {
	var message = models.Message{
		Sender:    createData.Sender,
		Receiver:  createData.Receiver,
		Content:   createData.Content,
		Timestamp: time.Now(),
	}

	if !r.ValidUser(createData.Sender) {
		return nil, errors.New("Invalid sender")
	}

	if !r.ValidUser(createData.Receiver) {
		return nil, errors.New("Invalid receiver")
	}

	if createData.Sender == createData.Receiver {
		return nil, errors.New("Can't send message to self")
	}

	if !r.ValidContent(createData.Content) {
		return nil, errors.New("Content empty or too long")
	}

	result := r.GetDB().Create(&message)

	if result.Error != nil {
		return nil, result.Error
	}

	return &message, nil

}

func (r *Repository) ReadBySender(readBySender IReadBySender) ([]models.Message, error) {
	var messagesArray []models.Message
	var messagesMap = make(map[string]interface{})

	messagesMap["sender"] = readBySender.Sender

	result := r.GetDB().Where(messagesMap).Find(&messagesArray)

	if result.Error != nil {
		return nil, result.Error
	}

	return messagesArray, nil
}

func (r *Repository) ReadByReceiver(readBy IReadByReceiver) ([]models.Message, error) {
	var messagesArray []models.Message
	var messageMap = make(map[string]interface{})

	messageMap["receiver"] = readBy.Receiver

	result := r.GetDB().Where(messageMap).Find(&messagesArray)

	if result.Error != nil {
		return nil, result.Error
	}

	return messagesArray, nil
}

func (r *Repository) ReadChat(readChat IReadChat) ([]models.Message, error) {
	var messagesArray []models.Message
	var messageMap = make(map[string]interface{})

	messageMap["sender"] = readChat.Sender
	messageMap["receiver"] = readChat.Receiver

	result := r.GetDB().Where(messageMap).Find(&messagesArray)

	if result.Error != nil {
		return nil, result.Error
	}

	return messagesArray, nil
}

func (r *Repository) Update(updateData IUpdate) (*models.Message, error) {
	var message = models.Message{ID: updateData.ID}
	verifyExistence := r.GetDB().First(&message)

	if verifyExistence.Error != nil {
		return nil, verifyExistence.Error
	}
    
    if !r.ValidContent(updateData.Content) {
		return nil, errors.New("Content empty or too long")
	}

	message.Content = updateData.Content
	saveResult := r.GetDB().Save(&message)

	if saveResult.Error != nil {
		return nil, saveResult.Error
	}

	return &message, nil
}

func (r *Repository) Delete(deleteData IDelete) (*models.Message, error) {
	var message = models.Message{ID: deleteData.ID}

	verifyExistence := r.GetDB().First(&message)

	if verifyExistence.Error != nil {
		return nil, verifyExistence.Error
	}

	result := r.GetDB().Delete(&message)

	if result.Error != nil {
		return nil, result.Error
	}

	return &message, nil
}
