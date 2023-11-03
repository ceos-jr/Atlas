package projectrepotest

import (
	"orb-api/config"
	"orb-api/models"
	repository "orb-api/repositories"
	//"orb-api/repositories/user"
	"github.com/stretchr/testify/suite"
)
type TestSuite struct {
	suite.Suite
	Repo      *repository.Repository
	MockProject []models.Project
}

func (suite *TestSuite) SetupSuite() {
	repo, setupError := config.SetupDB("../../.env")

	if setupError != nil {
		panic(setupError)
	}

	suite.Repo = repo
	suite.MockProject = make([]models.Project, 2)
	//suite.MockSector = make([]models.Sector, 2)
	suite.SetupMocks()
}

func (suite *TestSuite) SetupMocks() {

}