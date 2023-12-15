package taskservicetest

import (
	"fmt"
	"orb-api/config"
	"orb-api/models"
	"orb-api/services/task"
	"time"

	repository "orb-api/repositories"
	taskrepo "orb-api/repositories/task"
	userrepo "orb-api/repositories/user"

	"github.com/stretchr/testify/suite"
)

type TestSuit struct {
	suite.Suite
	Repo        *repository.Repository
	TaskService *task.Service
	MockUsers   []models.User
	MockTasks   []models.Task
}

// SetupSuite Executed before all tests
func (suite *TestSuit) SetupSuite() {
	repository, setupError := config.SetupDB("../../.env")

	if setupError != nil {
		panic(setupError)
	}

	suite.Repo = repository
	suite.TaskService = task.SetupTaskService(&repository.Task)
	suite.MockUsers = make([]models.User, 3)
	suite.MockTasks = make([]models.Task, 3)
	suite.SetupMocks()
}

// setting up the mock task
func (suite *TestSuit) SetupMocks() {
	for i := 0; i < 3; i++ {
		NewUser, createErr := suite.Repo.User.Create(userrepo.ICreate{
			Name:     fmt.Sprintf("Gabrigas %v", i+1),
			Email:    fmt.Sprintf("example0%v@example.com", i+1),
			Password: "gabrigas123",
			Status:   2,
		})

		if createErr != nil {
			panic(createErr)
		}

		suite.MockUsers[i] = *NewUser
	}
	
	deadline := time.Date(2077, 4, 12, 12, 0, 0, 0, time.Local)

	newTask, createErr := suite.Repo.Task.Create(taskrepo.ICreate{
		Description: "Uma tarefa",
		AssignedTo:  suite.MockUsers[0].ID,
		CreatedBy:   suite.MockUsers[1].ID,
		Status:      2,
		Deadline:    deadline,
	})

	if createErr != nil {
		panic(createErr)
	}

	
	deadline2 := time.Date(2077, 4, 12, 12, 0, 0, 0, time.Local)
	
	newTask2, createErr := suite.Repo.Task.Create(taskrepo.ICreate{
		Description: "Uma tarefa2",
		AssignedTo:  suite.MockUsers[2].ID,
		CreatedBy:   suite.MockUsers[1].ID,
		Status:      2,
		Deadline:    deadline2,
	})

	if createErr != nil {
		panic(createErr)
	}

	deadline3 := time.Date(2050, 4, 12, 12, 0, 0, 0, time.Local)

	newTask3, createErr := suite.Repo.Task.Create(taskrepo.ICreate{
		Description: "Uma tarefa3",
		AssignedTo:  suite.MockUsers[2].ID,
		CreatedBy:   suite.MockUsers[1].ID,
		Status:      2,
		Deadline:    deadline3,
	})

	if createErr != nil {
		panic(createErr)
	}


	suite.MockTasks[0] = *newTask
	suite.MockTasks[1] = *newTask2
	suite.MockTasks[2] = *newTask3
	
}

// TearDownSuite Executed after all tests
func (suite *TestSuit) TearDownSuite() {
	for index := range suite.MockUsers {
		_, deleteErr := suite.Repo.User.Delete(userrepo.IDelete{
			ID: suite.MockUsers[index].ID,
		})

		if deleteErr != nil {
			panic(deleteErr)
		}
	}

	for index := range suite.MockTasks {
		_, deleteErr := suite.Repo.Task.Delete(taskrepo.IDelete{
			ID: suite.MockTasks[index].ID,
		})

		if deleteErr != nil {
			panic(deleteErr)
		}
	}
}
