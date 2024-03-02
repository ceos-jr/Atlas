package secrepotest

import (
	"fmt"
	"orb-api/config"
	"orb-api/models"
	repository "orb-api/repositories"
	"orb-api/repositories/sector"
	"orb-api/repositories/user"

	"github.com/stretchr/testify/suite"
)

type TestSuite struct {
	suite.Suite
	Repo      *repository.Repository
	MockSector []models.Sector
	MockUsers  []models.User
}


func (suite *TestSuite) SetupSuite() {

	repo, setupError := config.SetupDB("../../.env")

	if setupError != nil {
		panic(setupError)
	}

	suite.Repo = repo
	suite.MockSector = make([]models.Sector, 6)
	suite.MockUsers = make([]models.User, 2)
	suite.SetupMocks()
}

func (suite *TestSuite) SetupMocks() {
	NewUser, createErr := suite.Repo.User.Create(user.ICreate{
		Name:     "User 01",
		Email:    "user01@example.com",
		Password: "12345678",
		Status:   1,
	})

	if createErr != nil {
		panic(createErr)
	}

	suite.MockUsers[0] = *NewUser

	NewUser2, createErr := suite.Repo.User.Create(user.ICreate{
		Name:     "User 02",
		Email:    "user02@example.com",
		Password: "12345678",
		Status:   1,
	})

	if createErr != nil {
		panic(createErr)
	}

	suite.MockUsers[1] = *NewUser2

	for i := 0; i < 5; i++{
		NewSector, createErr := suite.Repo.Sector.Create(sector.ICreate{
			Name: fmt.Sprintf("Sector %v", i+1),
			Description: fmt.Sprintf("Sector description %v", i+1),
			AdmID: suite.MockUsers[0].ID,
		})

		if createErr != nil {
			panic(createErr)
		}

		suite.MockSector[i] = *NewSector
	}
}

func (suite *TestSuite) TearDownSuite() {
	for index := range suite.MockUsers {
		_, deleteErr := suite.Repo.User.Delete(user.IDelete{
			ID: suite.MockUsers[index].ID,
		})

		if deleteErr != nil {
			panic(deleteErr)
		}
	}

	for i := range suite.MockSector {
		_, deleteErr := suite.Repo.Sector.Delete(sector.IDelete{
			ID: suite.MockSector[i].ID,
		})

		if deleteErr != nil {
			panic(deleteErr)
		}
	}
}
