package controllers

import (
	"orb-api/config"
	"orb-api/controllers/user"
	"orb-api/controllers/task"
	"orb-api/controllers/project"
	"orb-api/services"
)

type Controllers struct {
	User user.BaseHandler
	Task task.BaseHandler
	Project project.BaseHandler
}

func SetupControllers(service *services.Service) *Controllers {
	validator := config.NewValidator()

	return &Controllers{
		User: *user.NewBaseHandler(&service.User, validator),
		Task: *task.NewBaseHandler(&service.Task, validator),
		Project: *project.NewBaseHandler(&service.Project, validator),
	}
}
