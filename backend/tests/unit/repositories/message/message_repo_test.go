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

	suite.Equal("Content empty or too long",
		createErr.Error(), "Expected to have an error",
	)

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

func (suite *MessageRepoTestSuite) TestUpdateMessage() {
	content := "edited message"

	updatedMessage, updateError := suite.Repo.Message.Update(message.IUpdate{
		ID:      suite.MockMessages[1].ID,
		Content: content,
	})

	suite.Nil(updateError, "Update error must be nil")
	suite.Equal(content, updatedMessage.Content, "Content do not match")
	suite.Equal(suite.MockUsers[0].ID, updatedMessage.Sender, "Sender do not match")
	suite.Equal(suite.MockUsers[1].ID, updatedMessage.Receiver, "Receiver do not match")
}

func (suite *MessageRepoTestSuite) TestUpdateMessageErr() {
	invalidID := uint(9999)
	invalidContent := ""

	_, updateError := suite.Repo.Message.Update(message.IUpdate{
		ID:      invalidID,
		Content: "This is a simple test",
	})

	suite.Equal("record not found", updateError.Error(),
		"Invalid ID it should return an error",
	)

	_, updateError = suite.Repo.Message.Update(message.IUpdate{
		ID:      suite.MockMessages[0].ID,
		Content: invalidContent,
	})

	suite.Equal("Content empty or too long", updateError.Error(),
		"Invalid content it should return an error",
	)
}

func (suite *MessageRepoTestSuite) TestDeleteMessage() {
	newMessage, _ := suite.Repo.Message.Create(message.ICreate{
		Content:  "Delete message test",
		Sender:   suite.MockUsers[0].ID,
		Receiver: suite.MockUsers[1].ID,
	})

	deletedMessage, deleteErr := suite.Repo.Message.Delete(message.IDelete{
		ID: newMessage.ID,
	})

	suite.Nil(deleteErr, "Delete error must be nil")
	suite.Equal(newMessage.ID, deletedMessage.ID, "Expected to have the same ID")
}

func TestMessageRepository(test *testing.T) {
	suite.Run(test, new(MessageRepoTestSuite))
}
