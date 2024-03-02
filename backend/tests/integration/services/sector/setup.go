package sectorservicetest

import (
	"fmt"
	"orb-api/config"
	"orb-api/models"
	repository "orb-api/repositories"
	userrepo "orb-api/repositories/user"
	sectorrepo "orb-api/repositories/sector"


	"orb-api/services/sector"

	"github.com/stretchr/testify/suite"
)

type TestSuit struct {
	suite.Suite
	Repo          *repository.Repository
	SectorService *sector.Service
	MockUsers     []models.User
	MockSector    []models.Sector
}

func (suite *TestSuit) SetupSuite() {
	repository, setupError := config.SetupDB("../../.env")

	if setupError != nil {
		panic(setupError)
	}

	suite.Repo = repository
	suite.SectorService = sector.SetupSectorService(&repository.Sector)
	suite.MockUsers = make([]models.User, 3)
	suite.MockSector = make([]models.Sector, 4)
	suite.SetupMocks()
}

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

	for i := 0; i < 3; i++ {
		NewSector, createErr := suite.Repo.Sector.Create(sectorrepo.ICreate{
			Name:        fmt.Sprintf("Sector %v", i+1),
			Description: fmt.Sprintf("Description %v", i+1),
			AdmID:       suite.MockUsers[0].ID,
		})

		if createErr != nil {
			panic(createErr)
		}

		suite.MockSector[i] = *NewSector
		
	}
}



func (suite *TestSuit) TearDownSuite() {
	for index := range suite.MockUsers {
		_, deleteErr := suite.Repo.User.Delete(userrepo.IDelete{
			ID: suite.MockUsers[index].ID,
		})

		if deleteErr != nil {
			panic(deleteErr)
		}
	}

	for index := range suite.MockSector {
		_, deleteErr := suite.Repo.Sector.Delete(sectorrepo.IDelete{
			ID: suite.MockSector[index].ID,
		})

		if deleteErr != nil {
			panic(deleteErr)
		}
	}
}