package services

import (
	repository "orb-api/repositories"
	"orb-api/services/role"
	"orb-api/services/user"
	"orb-api/services/task"
)

type Service struct {
	User user.Service
	Role role.Service
	Task task.Service
}

func SetupServices(repository *repository.Repository) *Service {
	return &Service{
		User: *user.SetupService(&repository.User),
		Role: *role.Setup(&repository.Role),
		Task: *task.SetupTask(&repository.Task),
	}
}
