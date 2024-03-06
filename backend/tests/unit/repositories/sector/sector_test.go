package secrepotest

import (
	"orb-api/repositories/sector"
	"testing"

	"github.com/stretchr/testify/suite"
)

func (suite *TestSuite) TestCreateSector() {

	NewSector, createErr := suite.Repo.Sector.Create(sector.ICreate{
		Name:        "Sector 1",
		Description: "Sector description 5",
		AdmID:       suite.MockUsers[0].ID,
	})

	suite.Nil(createErr, "Create error must be nil")
	suite.Equal("Sector 1", NewSector.Name)
	suite.Equal("Sector description 5", NewSector.Description)
	suite.Equal(suite.MockUsers[0].ID, NewSector.AdmID)

	suite.MockSector[5] = *NewSector
}

func (suite *TestSuite) TestReadAllSectors() {
	sectors, readErr := suite.Repo.Sector.ReadAll()

	suite.Nil(readErr, "Read error must be nil")
	suite.Equal(7, len(sectors), "Expected to have seven sectors")
}

func (suite *TestSuite) TestReadByID() {
	sector, readErr := suite.Repo.Sector.ReadBy(sector.IReadBy{
		ID: &suite.MockSector[0].ID,
	})

	suite.Nil(readErr, "Read error must be nil")
	suite.Equal(suite.MockSector[0].ID, sector[0].ID, "Expected to have the same ID")
}

func (suite *TestSuite) TestReadSectorByErr() {
	_, readErr := suite.Repo.Sector.ReadBy(sector.IReadBy{})

	suite.Equal(readErr.Error(), "No fields to read" ,"Expected to have an error")
}


func (suite *TestSuite) TestUpdateSector() {
	name := "Sector 1 updated"
	description := "Sector description 1 updated"
	admID := suite.MockUsers[1].ID

	sector, updateErr := suite.Repo.Sector.Update(sector.IUpdate{
		ID:          suite.MockSector[1].ID,
		Name:        &name,
		Description: &description,
		AdmID:       &admID,
	})

	suite.Nil(updateErr, "Update error must be nil")
	suite.Equal("Sector 1 updated", sector.Name)
	suite.Equal("Sector description 1 updated", sector.Description)
	suite.Equal(suite.MockUsers[1].ID, sector.AdmID)

	suite.MockSector[1] = *sector
}

func (suite *TestSuite) TestUpdateSectorErr() {
	_, updateErr := suite.Repo.Sector.Update(sector.IUpdate{})

	suite.Equal(updateErr.Error(), "No fields to update", "Expected to have an error")

}

func (suite *TestSuite) TestDeleteSector() {
	newSector, _ := suite.Repo.Sector.Create(sector.ICreate{
		Name:        "New Sector",
		Description: "New Sector description",
		AdmID:       suite.MockUsers[0].ID,
	})

	deletedSector, deleteErr := suite.Repo.Sector.Delete(sector.IDelete{
		ID: newSector.ID,
	})

	suite.Nil(deleteErr, "Delete error must be nil")
	suite.Equal(newSector.ID, deletedSector.ID, "Expected to have the same id")

}


func TestSectorRepository(test *testing.T) {
	suite.Run(test, new(TestSuite))
}
