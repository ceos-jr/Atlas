package controllers

import (
	"orb-api/config"
	"orb-api/controllers/project"
	"orb-api/controllers/sector"
	"orb-api/controllers/task"
	"orb-api/controllers/user"
	"orb-api/services"
)

type Controllers struct {
	User    user.BaseHandler
	Task    task.BaseHandler
	Project project.BaseHandler
	Sector  sector.BaseHandler
}

func SetupControllers(service *services.Service) *Controllers {
	validator := config.NewValidator()

	return &Controllers{
		User:    *user.NewBaseHandler(&service.User, validator),
		Task:    *task.NewBaseHandler(&service.Task, validator),
		Project: *project.NewBaseHandler(&service.Project, validator),
		Sector:  *sector.NewBaseHandler(&service.Sector, validator),
	}
}
