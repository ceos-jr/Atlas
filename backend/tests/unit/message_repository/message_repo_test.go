package messagerepotest

import (
	"orb-api/repositories/message"
	"testing"

	"github.com/stretchr/testify/suite"
)

func (suite *MessageRepoTestSuite) TestCreateMessage() {

	message, createErr := suite.Repo.Message.Create(message.ICreate{
		Sender:   suite.MockUsers[0].ID,
		Receiver: suite.MockUsers[1].ID,
		Content:  "This is a message",
	})

	suite.Nil(createErr, "Create error must be nil")
	suite.Equal(suite.MockUsers[0].ID, message.Sender)
	suite.Equal(suite.MockUsers[1].ID, message.Receiver)
	suite.Equal("This is a message", message.Content)

}

func (suite *MessageRepoTestSuite) TestCreateMessagErr() {
	var sender uint
	var receiver uint
	var content string

	var validSender uint = 1
	var invalidSender uint = 999
	var validReceiver uint = 2
	var invalidReceiver uint = 999
	validContent := "This is a valid message"
	invalidContent := ""

	for index := 0; index < 8; index++ {

		if index&0b001 != 0 {
			sender = validSender
		} else {
			sender = invalidSender
		}

		if index&0b010 != 0 {
			receiver = validReceiver
		} else {
			receiver = invalidReceiver
		}

		if index&0b100 != 0 {
			content = validContent
		} else {
			content = invalidContent
		}

		_, createErr := suite.Repo.Message.Create(message.ICreate{
			Sender:   sender,
			Receiver: receiver,
			Content:  content,
		})

		suite.NotNil(createErr, "Invalid message")
	}

	_, createErr := suite.Repo.Message.Create(message.ICreate{
		Sender:   sender,
		Receiver: sender,
		Content:  "Olá, tudo bem?",
	})
	suite.Equal(createErr, "Can't send message to self")

}

func (suite *MessageRepoTestSuite) TestReadMessageBySender() {
	messages, readErr := suite.Repo.Message.ReadBySender(message.IReadBySender{
		Sender: suite.MockUsers[0].ID,
	})

	suite.Nil(readErr, "Read Error must be nil")
	suite.Equal(suite.MockUsers[0].ID, messages[0].Sender, "IDs must match")

}

func (suite *MessageRepoTestSuite) TestReadMessageByReceiver() {
	messages, readErr := suite.Repo.Message.ReadByReceiver(message.IReadByReceiver{
		Receiver: suite.MockUsers[0].ID,
	})

	suite.Nil(readErr, "Read Error must be nil")
	suite.Equal(suite.MockUsers[0].ID, messages[0].Sender, "IDs must match")
}

//func (suite *MessageRepoTestSuite) TestReadChat() {
//	messages, readErr := suite.R
//}

func TestMessageRepository(test *testing.T) {
	suite.Run(test, new(MessageRepoTestSuite))
}
