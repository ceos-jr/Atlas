package messagerepotest

import (
	"fmt"
	"orb-api/config"
	"orb-api/models"
	repository "orb-api/repositories"
	"orb-api/repositories/message"
	"orb-api/repositories/user"

	"github.com/stretchr/testify/suite"
)

type MessageRepoTestSuite struct {
	suite.Suite
	Repo         *repository.Repository
	MockUsers    []models.User
	MockMessages []models.Message
}

func (suite *MessageRepoTestSuite) SetupSuite() {
	repo, setupError := config.SetupDB("../../.env")

	if setupError != nil {
		panic(setupError)
	}

	suite.Repo = repo
	suite.MockUsers = make([]models.User, 2)
	suite.MockMessages = make([]models.Message, 2)
	suite.SetupMocks()
}

func (suite *MessageRepoTestSuite) SetupMocks() {
	for index := 0; index < 2; index++ {
		user, createErr := suite.Repo.User.Create(user.ICreate{
			Name:     fmt.Sprintf("Mahmoud Mahmed %v", index+1),
			Email:    fmt.Sprintf("mahmoud123@gmail.com %v", index+1),
			Password: "123456789",
			Status:   2,
		})

		if createErr != nil {
			panic(createErr)
		}

		suite.MockUsers[index] = *user
	}

	message, createErr := suite.Repo.Message.Create(message.ICreate{
		Sender:   suite.MockUsers[0].ID,
		Receiver: suite.MockUsers[1].ID,
		Content:  fmt.Sprintf("Do that 0"),
	})

	if createErr != nil {
		panic(createErr)
	}

	suite.MockMessages[0] = *message
}

func (suite *MessageRepoTestSuite) TearDownSuite() {
	for index := range suite.MockUsers {
		_, deleteErr := suite.Repo.User.Delete(user.IDelete{
			ID: suite.MockUsers[index].ID,
		})

		if deleteErr != nil {
			panic(deleteErr)
		}
	}

	for index := range suite.MockMessages {
		_, deleteErr := suite.Repo.Message.Delete(message.IDelete{
			ID: suite.MockMessages[index].ID,
		})

		if deleteErr != nil {
			panic(deleteErr)
		}
	}
}
