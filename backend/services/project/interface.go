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
		CreateProject(createData project.ICreate) (*models.Project, error)
	}
)
