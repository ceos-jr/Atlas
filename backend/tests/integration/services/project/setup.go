
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
	userrepo "orb-api/repositories/user"

	"github.com/stretchr/testify/suite"
)

type TestSuit struct {
	suite.Suite
	Repo        *repository.Repository
	ProjectService *project.Service
	MockUsers []models.User
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
  	suite.ProjectService = project.SetupProjectService(&repository.Project, &repository.UserProject, &repository.TaskProject, &repository.Task, &repository.User)
	suite.MockProjects = make([]models.Project, 6)
	suite.MockTasks = make([]models.Task, 2)
	suite.MockUsers = make([]models.User, 4)
	suite.SetupMocks()
}

// setting up the mock task
func (suite *TestSuit) SetupMocks() {

	NewUser, createuserErr := suite.Repo.User.Create(userrepo.ICreate{
		Name:     fmt.Sprintf("Gabrigas %v", 1),
		Email:    fmt.Sprintf("example0%v@example.com", 1),
		Password: "gabrigas123",
		Status:   2,
	})

	if createuserErr != nil {
		panic(createuserErr)
	}

	suite.MockUsers[0] = *NewUser

	NewUser2, createuserErr2 := suite.Repo.User.Create(userrepo.ICreate{
		Name:     fmt.Sprintf("Gabrigas %v", 2),
		Email:    fmt.Sprintf("example0%v@example.com", 2),
		Password: "gabrigas123",
		Status:   2,
	})

	if createuserErr2 != nil {
		panic(createuserErr2)
	}

	suite.MockUsers[1] = *NewUser2

	NewUser3, createuserErr3 := suite.Repo.User.Create(userrepo.ICreate{
		Name:     fmt.Sprintf("Gabrigas %v", 3),
		Email:    fmt.Sprintf("example0%v@example.com", 3),
		Password: "gabrigas123",
		Status:   2,
	})

	if createuserErr3 != nil {
		panic(createuserErr3)
	}

	suite.MockUsers[2] = *NewUser3

	NewUser4, createuserErr4 := suite.Repo.User.Create(userrepo.ICreate{
		Name:     fmt.Sprintf("Gabrigas %v", 4),
		Email:    fmt.Sprintf("example0%v@example.com", 4),
		Password: "gabrigas123",
		Status:   2,
	})

	if createuserErr4 != nil {
		panic(createuserErr4)
	}

	suite.MockUsers[3] = *NewUser4

	NewTask, createErr := suite.Repo.Task.Create(taskrepo.ICreate{
		Description: "Uma tarefa",
		AssignedTo:  suite.MockUsers[1].ID,
		CreatedBy:   suite.MockUsers[0].ID,
		Status:      2,
		Deadline:    time.Date(2077, time.December, 16, 12, 0, 0, 0, time.Local),
	})

	if createErr != nil {
		panic(createErr)
	}

	suite.MockTasks[0] = *NewTask

	NewTask2, createErr := suite.Repo.Task.Create(taskrepo.ICreate{
		Description: "Uma tarefa",
		AssignedTo: suite.MockUsers[1].ID,
		CreatedBy:	suite.MockUsers[0].ID,
		Status:		2, 
		Deadline:	time.Date(2077, time.December, 15, 12, 0, 0, 0, time.Local),
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


	NewProject2, createErr2 := suite.Repo.Project.Create(projectrepo.ICreate{
		Name: fmt.Sprintf("Projeto2"),
		Sector: 1,
		AdmID: 1,
	})

	if createErr2 != nil {
		panic(createErr2)
	}

	suite.MockProjects[1] = *NewProject2

	NewProject3, createErr3 := suite.Repo.Project.Create(projectrepo.ICreate{
		Name: fmt.Sprintf("Projeto3"),
		Sector: 2,
		AdmID: 2,
	})

	if createErr3 != nil {
		panic(createErr3)
	}

	suite.MockProjects[2] = *NewProject3

	NewProject4, createErr4 := suite.Repo.Project.Create(projectrepo.ICreate{
		Name: fmt.Sprintf("Projeto4"),
		Sector: 2,
		AdmID: 2,
	})

	if createErr4 != nil {
		panic(createErr3)
	}

	suite.MockProjects[3] = *NewProject4

	NewProject5, createErr5 := suite.Repo.Project.Create(projectrepo.ICreate{
		Name: fmt.Sprintf("Projeto5"),
		Sector: 2,
		AdmID: 2,
	})

	if createErr5 != nil {
		panic(createErr5)
	}

	suite.MockProjects[4] = *NewProject5

	NewProject6, createErr6 := suite.Repo.Project.Create(projectrepo.ICreate{
		Name: fmt.Sprintf("Projeto6"),
		Sector: 2,
		AdmID: 2,
	})

	if createErr6 != nil {
		panic(createErr6)
	}

	suite.MockProjects[5] = *NewProject6


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
			ID: suite.MockProjects[index].ID,
		})

		if deleteErr != nil {
			panic(deleteErr)
		}
	}

	for index := range suite.MockUsers {
		_, deleteErr := suite.Repo.User.Delete(userrepo.IDelete{
			ID: suite.MockUsers[index].ID,
		})

		if deleteErr != nil {
			panic(deleteErr)
		}
	}
}