package sector

import (
	"errors"
	"orb-api/models"
	"orb-api/repositories/sector"
)

func SetupSectorService(repo *sector.Repository) *Service {
	return &Service{
		SectorRepo: repo,
	}
}

func (s *Service) CreateSector(name string, description string, admid uint) (*models.Sector, error) {
	if description == "" || name == ""{
		return nil, errors.New("description or name cannot be empty")
	}

	sectorArray, readErr := s.SectorRepo.ReadBy(sector.IReadBy{
		Name: &name,
	})

	if readErr != nil {
		return nil, readErr
	}

	if len(sectorArray) != 0 {
		return nil, errors.New("This sector name is already being used")
	}
	

	createdSector, createErr := s.SectorRepo.Create(sector.ICreate{
		Name:        name,
		Description: description,
		AdmID: 	     admid,
	})

	if createErr != nil {
		return nil, createErr
	}

	return createdSector, nil
}


