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
	if description == ""{
		return nil, errors.New("description cannot be empty")
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


