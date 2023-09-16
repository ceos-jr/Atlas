package services

import (
	"orb-api/repositories"
	"orb-api/services/user"
)

type Service struct {
	User user.Service
}

func SetupServices(repository *repository.Repository) *Service {
	return &Service{
		User: *user.SetupService(&repository.User),
	}
}
