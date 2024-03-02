package services

import (
	repository "orb-api/repositories"
	"orb-api/services/project"
	"orb-api/services/role"
	"orb-api/services/sector"
	"orb-api/services/task"
	"orb-api/services/user"
	"orb-api/services/user_role"
)

type Service struct {
	User     user.Service
	Role     role.Service
	Task     task.Service
	UserRole user_role.Service
	Project  project.Service
	Sector   sector.Service
}

func SetupServices(repository *repository.Repository) *Service {
	return &Service{
		User:     *user.SetupService(&repository.User, &repository.UserProject, &repository.Project),
		Role:     *role.Setup(&repository.Role),
		UserRole: *user_role.SetupService(&repository.User, &repository.Role, &repository.UserRole),
		Task:     *task.SetupTaskService(&repository.Task),
		Project:  *project.SetupProjectService(&repository.Project, &repository.UserProject, &repository.TaskProject, &repository.Task, &repository.User),
		Sector:   *sector.SetupSectorService(&repository.Sector),
	}
}
