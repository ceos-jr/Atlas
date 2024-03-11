package controllers

import (
	"orb-api/config"
	"orb-api/controllers/project"
	"orb-api/controllers/role"
	"orb-api/controllers/task"
	"orb-api/controllers/user"
	"orb-api/controllers/user_role"
	"orb-api/services"
)

type Controllers struct {
	User user.BaseHandler
	Task task.BaseHandler
	Project project.BaseHandler
	Role role.BaseHandler
	User_Role user_role.BaseHandler
}

func SetupControllers(service *services.Service) *Controllers {
	validator := config.NewValidator()

	return &Controllers{
		User: *user.NewBaseHandler(&service.User, validator),
		Task: *task.NewBaseHandler(&service.Task, validator),
		Project: *project.NewBaseHandler(&service.Project, validator),
		Role:  *role.NewBaseHandler(&service.Role, validator),
		User_Role: *user_role.NewBaseHandler(&service.UserRole, validator),
	}
}
