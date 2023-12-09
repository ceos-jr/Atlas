
package projectservicetest

import (
	"time"
	"fmt"
	"orb-api/config"
	"orb-api/models"
	"orb-api/services/project"

	repository "orb-api/repositories"
	taskrepo	"orb-api/repositories/task"
	projectrepo	"orb-api/repositories/project"

	"github.com/stretchr/testify/suite"
)

type TestSuit struct {
	suite.Suite
	Repo        *repository.Repository
	ProjectService *project.Service
	MockProjects []models.Project
	MockTasks	[]models.Task
	MockTasksProjects	[]models.TasksProject
}

// SetupSuite Executed before all tests
func (suite *TestSuit) SetupSuite() {
	repository, setupError := config.SetupDB("../../.env")

	if setupError != nil {
		panic(setupError)
	}

	suite.Repo = repository
  	suite.ProjectService = project.SetupProjectService(&repository.Project, &repository.UserProject, &repository.TaskProject, &repository.Task)
	suite.MockProjects = make([]models.Project, 1)
	suite.MockTasks = make([]models.Task, 2)
	suite.SetupMocks()
}

// setting up the mock task
func (suite *TestSuit) SetupMocks() {

	NewTask, createErr := suite.Repo.Task.Create(taskrepo.ICreate{
		Description: "Uma tarefa",
		AssignedTo:  1,
		CreatedBy:   2,
		Status:      2,
		Deadline:    time.Date(2023, time.December, 16, 12, 0, 0, 0, time.Local),
	})

	if createErr != nil {
		panic(createErr)
	}

	suite.MockTasks[0] = *NewTask

	NewTask2, createErr := suite.Repo.Task.Create(taskrepo.ICreate{
		Description: "Uma tarefa",
		AssignedTo: 1,
		CreatedBy:	2,
		Status:		2, 
		Deadline:	time.Date(2023, time.December, 15, 12, 0, 0, 0, time.Local),
	})

	if createErr != nil {
		panic(createErr)
	}

	suite.MockTasks[1] = *NewTask2

	NewProject, createErr := suite.Repo.Project.Create(projectrepo.ICreate{
		Name:	fmt.Sprintf("Projeto"),
		Sector:	1,
		AdmID:	1,
	})

	if createErr != nil {
		panic(createErr)
	}

	suite.MockProjects[0] = *NewProject
}

// TearDownSuite Executed after all tests
func (suite *TestSuit) TearDownSuite() {
	for index := range suite.MockTasks {
		_, deleteErr := suite.Repo.Task.Delete(taskrepo.IDelete{
			ID: suite.MockTasks[index].ID,
		})

		if deleteErr != nil {
			panic(deleteErr)
		}
	}

	for index := range suite.MockProjects {
		_, deleteErr := suite.Repo.Project.Delete(projectrepo.IDelete{
			ID: suite.MockTasks[index].ID,
		})

		if deleteErr != nil {
			panic(deleteErr)
		}
	}
}
