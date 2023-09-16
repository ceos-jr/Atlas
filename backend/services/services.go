package services

import (
	"orb-api/repositories"
	"orb-api/services/role"
	"orb-api/services/user"
)

type Service struct {
	User user.Service
	Role role.Service
}

func SetupServices(repository *repository.Repository) *Service {
	return &Service{
		User: *user.SetupService(&repository.User),
		Role: *role.SetupService(&repository.Role),
	}
}
