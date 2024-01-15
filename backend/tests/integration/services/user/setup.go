package userservicetest

import (
	"orb-api/config"
	"orb-api/models"
	"orb-api/services/user"

	userrepo "orb-api/repositories/user"
	projectrepo "orb-api/repositories/project"
	userprojectrepo "orb-api/repositories/userproject"

	"github.com/stretchr/testify/suite"
	"fmt"
)

type TestSuit struct {
	suite.Suite
	Service   *user.Service
	MockUsers []models.User
	MockProject models.Project
	MockUserProject models.UsersProject
}

// SetupSuite Executed before all tests
func (suite *TestSuit) SetupSuite() {
	repositories, setupError := config.SetupDB("../../.env")

	if setupError != nil {
		panic(setupError)
	}

	suite.Service = user.SetupService(&repositories.User, &repositories.UserProject, &repositories.Project)
	suite.MockUsers = make([]models.User, 6)
	suite.SetupMocks()
}

func (suite *TestSuit) SetupMocks() {
	newUser, createErr := suite.Service.UserRepo.Create(userrepo.ICreate{
		Name:     "Gabrigas",
		Email:    "gabrigas@example.com",
		Password: "mostBeautiful",
		Status:   1,
	})

	if createErr != nil {
		panic(createErr)
	}

	suite.MockUsers[0] = *newUser

	newUser2, createErr2 := suite.Service.UserRepo.Create(userrepo.ICreate{
		Name:     "Gabrigas2",
		Email:    "gabrigas2@example.com",
		Password: "mostBeautiful",
		Status:   2,
	})

	if createErr2 != nil {
		panic(createErr)
	}

	suite.MockUsers[5] = *newUser2

	newUser3, createErr3 := suite.Service.UserRepo.Create(userrepo.ICreate{
		Name:     "Gabrigas3",
		Email:    "gabrigas3@example.com",
		Password: "mostBeautiful",
		Status:   1,
	})

	if createErr3 != nil {
		panic(createErr)
	}

	suite.MockUsers[2] = *newUser3

	newUser4, createErr4 := suite.Service.UserRepo.Create(userrepo.ICreate{
		Name:     "Gabrigas4",
		Email:    "gabrigas4@example.com",
		Password: "mostBeautiful",
		Status:   2,
	})

	if createErr4 != nil {
		panic(createErr)
	}

	suite.MockUsers[3] = *newUser4

	newUser5, createErr5 := suite.Service.UserRepo.Create(userrepo.ICreate{
		Name:     "Gabrigas5",
		Email:    "gabrigas5@example.com",
		Password: "mostBeautiful",
		Status:   2,
	})

	if createErr5 != nil {
		panic(createErr)
	}

	suite.MockUsers[4] = *newUser5
	
	NewProject, createErr4 := suite.Service.ProjectRepo.Create(projectrepo.ICreate{
		Name:	fmt.Sprintf("Projeto"),
		Sector:	1,
		AdmID:	1,
	})

	if createErr4 != nil {
		panic(createErr4)
	}

	suite.MockProject = *NewProject

	newUserProject, createErr5 := suite.Service.UserProjectRepo.Create(userprojectrepo.ICreate{
		UserID:	suite.MockUsers[0].ID,
		ProjectID: NewProject.ID,
	})

	if createErr5 != nil{
		panic(createErr5)
	}

	suite.MockUserProject = *newUserProject
}

// TearDownSuite Executed after all tests
func (suite *TestSuit) TearDownSuite() {
	for index := range suite.MockUsers {
		_, deleteErr := suite.Service.UserRepo.Delete(userrepo.IDelete{
			ID: suite.MockUsers[index].ID,
		})

		if deleteErr != nil {
			panic(deleteErr)
		}
	}

	_, deleteErr2 := suite.Service.ProjectRepo.Delete(projectrepo.IDelete{
		ID: suite.MockProject.ID,
	})

	if deleteErr2 != nil {
		panic(deleteErr2)
	}

	_, deleteErr3 := suite.Service.UserProjectRepo.Delete(userprojectrepo.IDelete{
		ID: suite.MockUserProject.ID,
	})

	if deleteErr3 != nil{
		panic(deleteErr3)
	}
}
