package taskrepotest

import (
	"fmt"
	"orb-api/config"
	"orb-api/models"
	repository "orb-api/repositories"
	"orb-api/repositories/task"
	"orb-api/repositories/user"
	"time"

	"github.com/stretchr/testify/suite"
)

type TaskRepoTestSuite struct {
	suite.Suite
	Repo      *repository.Repository
	MockUsers []models.User
	MockTasks []models.Task
}

// Executed before all tests
func (suite *TaskRepoTestSuite) SetupSuite() {
	repo, setupError := config.SetupDB("../.env")

	if setupError != nil {
		panic(setupError)
	}

	suite.Repo = repo
	suite.MockUsers = make([]models.User, 2)
	suite.MockTasks = make([]models.Task, 2)
	suite.SetupMocks()
}

// Executed after all tests
func (suite *TaskRepoTestSuite) TearDownSuite() {
	for index := range suite.MockUsers {
		_, deleteErr := suite.Repo.User.Delete(user.IDelete{
			ID: suite.MockUsers[index].ID,
		})

		if deleteErr != nil {
			panic(deleteErr)
		}
	}

	for index := range suite.MockTasks {
		_, deleteErr := suite.Repo.Task.Delete(task.IDelete{
			ID: suite.MockTasks[index].ID,
		})

		if deleteErr != nil {
			panic(deleteErr)
		}
	}
}

func (suite *TaskRepoTestSuite) SetupMocks() {
	for index := 0; index < 2; index++ {
		user, createErr := suite.Repo.User.Create(user.ICreate{
			Name:     fmt.Sprintf("Gabrigas %v", index+1),
			Email:    fmt.Sprintf("example0%v@example.com", index+1),
			Password: "gabrigas123",
			Status:   2,
		})

		if createErr != nil {
			panic(createErr)
		}

		suite.MockUsers[index] = *user
	}

	task, createErr := suite.Repo.Task.Create(task.ICreate{
		Description: "This is a mock task",
		CreatedBy:   suite.MockUsers[0].ID,
		AssignedTo:  suite.MockUsers[1].ID,
		Status:      2,
		Deadline:    time.Date(2077, 4, 12, 12, 0, 0, 0, time.UTC),
	})

	if createErr != nil {
		panic(createErr)
	}

	suite.MockTasks[0] = *task
}
