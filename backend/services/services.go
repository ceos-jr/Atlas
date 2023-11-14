package services

import (
	repository "orb-api/repositories"
	"orb-api/services/role"
	"orb-api/services/task"
	"orb-api/services/user"
	"orb-api/services/user_role"
)

type Service struct {
	User user.Service
	Role role.Service
	Task task.Service
	UserRole user_role.Service
}

func SetupServices(repository *repository.Repository) *Service {
	return &Service{
		User: *user.SetupService(&repository.User),
		Role: *role.Setup(&repository.Role),
		UserRole:	*user_role.SetupService(&repository.User, &repository.Role, &repository.UserRole),
		Task: *task.SetupTaskService(&repository.Task),
	}
}
