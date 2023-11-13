package project

import (
	"orb-api/models"
	"orb-api/repositories/project"
)

type (
	Service struct {
		ProjectRepo *project.Repository
	}

	Interface interface {
		CreateProject(name string, Sector uint, AdmID uint) (*models.Project, error)
	}
)
