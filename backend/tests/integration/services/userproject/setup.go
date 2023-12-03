package userprojectservicetest

import (
	"fmt"
	"orb-api/config"
	"orb-api/models"
	"orb-api/services/project"

	repository "orb-api/repositories"
	projectrepo "orb-api/repositories/project"
	userrepo "orb-api/repositories/user"

	"github.com/stretchr/testify/suite"
)

type TestSuit struct {
	suite.Suite
	Repo        *repository.Repository
	ProjectService *project.Service
	MockProjects   []models.Project
	MockUsers   []models.User
	MockUserProjects []models.UsersProject
}

// SetupSuite Executed before all tests
func (suite *TestSuit) SetupSuite() {
	repository, setupError := config.SetupDB("../../.env")

	if setupError != nil {
		panic(setupError)
	}

	suite.Repo = repository
  	suite.ProjectService = project.SetupProjectService(&repository.Project, &repository.UserProject)
	suite.MockUsers = make([]models.User, 1)
	suite.MockProjects = make([]models.Project, 1)
	suite.SetupMocks()
}

// setting up the mock task
func (suite *TestSuit) SetupMocks() {
	NewUser, createErr := suite.Repo.User.Create(userrepo.ICreate{
		Name:     fmt.Sprintf("Gabrigas %v", 1),
		Email:    fmt.Sprintf("example0%v@example.com", 1),
		Password: "gabrigas123",
		Status:   2,
	})

	if createErr != nil {
		panic(createErr)
	}

	suite.MockUsers[0] = *NewUser

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
	_, deleteErr := suite.Repo.User.Delete(userrepo.IDelete{
		ID: suite.MockUsers[0].ID,
	})

	if deleteErr != nil {
		panic(deleteErr)
	}

	_, deleteErr2 := suite.Repo.Project.Delete(projectrepo.IDelete{
		ID: suite.MockProjects[0].ID,
	})

	if deleteErr2 != nil {
		panic(deleteErr2)
	}
}