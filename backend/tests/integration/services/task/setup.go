package taskservicetest

import (
	"orb-api/config"
	"orb-api/models"
	"orb-api/services/task"
	"time"

	TaskRepo "orb-api/repositories/task"

	"github.com/stretchr/testify/suite"
)

type TestSuit struct {
	suite.Suite
	Service   *task.Service
	MockTasks []models.Task
}

// SetupSuite Executed before all tests
func (suite *TestSuit) SetupSuite() {
	repositories, setupError := config.SetupDB("../../.env")

	if setupError != nil {
		panic(setupError)
	}

	suite.Service = task.SetupService(&repositories.Task)
	suite.MockTasks = make([]models.Task, 2)
	suite.SetupMocks()
}

//setting up the mock task
func (suite *TestSuit) SetupMocks() {
	deadline := time.Date(2023, time.November, 15, 12, 0, 0, 0, time.UTC)

	newTask, createErr := suite.Service.TaskRepo.Create(taskrepo.ICreate{
		Description:	"Uma tarefa",
		AssignedTo:  	2,
		CreatedBy:   	1,
		Status:      	0,
		Deadline:    deadline,
	})

	if createErr != nil {
		panic(createErr)
	}

	suite.MockTasks[0] = *newTask
}

// TearDownSuite Executed after all tests
func (suite *TestSuit) TearDownSuite() {
	for index := range suite.MockTasks {
		_, deleteErr := suite.Service.TaskRepo.Delete(taskrepo.IDelete{
			ID: suite.MockTasks[index].ID,
		})

		if deleteErr != nil {
			panic(deleteErr)
		}
	}
}