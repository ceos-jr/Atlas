package sector

import (
	"orb-api/models"
	"orb-api/repositories/sector"
)

type (
	Service struct {
		SectorRepo *sector.Repository
	}

	Interface interface {
		CreateSector(name string, description string, admid uint) (*models.Sector, error)
	}
)
