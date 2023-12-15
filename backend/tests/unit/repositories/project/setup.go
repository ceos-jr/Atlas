package projectrepotest

import (
	"orb-api/config"
	"orb-api/models"
	repository "orb-api/repositories"
	"orb-api/repositories/project"
	"orb-api/repositories/user"
	"github.com/stretchr/testify/suite"
)
type ProjectTestSuite struct {
	suite.Suite
	Repo      *repository.Repository
	MockProject []models.Project
	MockUser []models.User
}

func (suite *ProjectTestSuite) SetupSuite() {
	repo, setupError := config.SetupDB("../../.env")

	if setupError != nil {
		panic(setupError)
	}

	suite.Repo = repo
	suite.MockProject = make([]models.Project, 2)
	suite.MockUser = make([]models.User, 1)
	//suite.MockSector = make([]models.Sector, 2)
	suite.SetupMocks()
}

func (suite *ProjectTestSuite) SetupMocks() {
	adm, createErr := suite.Repo.User.Create(user.ICreate{
		Name:     "User 01",
		Email:    "user01@example.com",
		Password: "12345678",
		Status:   1,
	})

	if createErr != nil {
		panic(createErr)
	}

	suite.MockUser[0] = *adm


	project, createErr := suite.Repo.Project.Create(project.ICreate{
		Name: "Projeto",
		AdmID: suite.MockUser[0].ID,
		Sector: 1,
	})

	if createErr != nil {
		panic(createErr)
	}

	suite.MockProject[0] = *project
}

func (suite *ProjectTestSuite) TearDownSuite() {
	for index := range suite.MockProject {
		_, deleteErr := suite.Repo.Project.Delete(project.IDelete{
			ID: suite.MockProject[index].ID,
		})

		if deleteErr != nil {
			panic(deleteErr)
		}
	}
}