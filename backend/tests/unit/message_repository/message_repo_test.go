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

func (suite *MessageRepoTestSuite) TestCreateMessageErr() {
	invalidSender := 9999
	invalidReceiver := 9999
	invalidContent := ""

	_, createErr := suite.Repo.Message.Create(message.ICreate{
		Sender:   uint(invalidSender),
		Receiver: suite.MockUsers[1].ID,
		Content:  "This is a message",
	})

	suite.Equal("Invalid sender", createErr.Error(), "Expected to have an error")

	_, createErr = suite.Repo.Message.Create(message.ICreate{
		Sender:   suite.MockUsers[0].ID,
		Receiver: uint(invalidReceiver),
		Content:  "This is a message",
	})

	suite.Equal("Invalid receiver", createErr.Error(), "Expected to have an error")

	_, createErr = suite.Repo.Message.Create(message.ICreate{
		Sender:   suite.MockUsers[0].ID,
		Receiver: suite.MockUsers[1].ID,
		Content:  invalidContent,
	})

	suite.Equal("Content empty or too long", createErr.Error(), "Expected to have an error")

	_, createErr = suite.Repo.Message.Create(message.ICreate{
		Sender:   suite.MockUsers[0].ID,
		Receiver: suite.MockUsers[0].ID,
		Content:  "This is a message",
	})

	suite.Equal("Can't send message to self", createErr.Error(),
		"Expected to have an error",
	)
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
		Receiver: suite.MockUsers[1].ID,
	})

	suite.Nil(readErr, "Read Error must be nil")
	suite.Equal(suite.MockUsers[1].ID, messages[0].Receiver, "IDs must match")
}

//func (suite *MessageRepoTestSuite) TestReadChat() {
//	messages, readErr := suite.R
//}

func TestMessageRepository(test *testing.T) {
	suite.Run(test, new(MessageRepoTestSuite))
}
