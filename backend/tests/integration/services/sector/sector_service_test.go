package sectorservicetest

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

func (suite *TestSuit) TestCreateSector() {
	NewSector, CreateErr := suite.SectorService.CreateSector("Sector 4", "Description 4", suite.MockUsers[0].ID)

	suite.Nil(CreateErr, "Create error must be nil")
	suite.Equal("Sector 4", NewSector.Name)

	suite.MockSector[3] = *NewSector
}

func (suite *TestSuit) TestCreateSectorErr() {
	_, CreateErr1 := suite.SectorService.CreateSector("", "Description 1", suite.MockUsers[0].ID)

	suite.Equal("description or name cannot be empty", CreateErr1.Error())

	_, CreateErr2 := suite.SectorService.CreateSector("Sector 1", "", suite.MockUsers[0].ID)

	suite.Equal("description or name cannot be empty", CreateErr2.Error())

	_, CreateErr3 := suite.SectorService.CreateSector("Sector 1", "Description 1", suite.MockUsers[0].ID)

	suite.Equal("This sector name is already being used", CreateErr3.Error())

}

func TestSectorRepository(t *testing.T) {
	suite.Run(t, new(TestSuit))
}
