package controllers

import (
	"orb-api/config"
	"orb-api/controllers/user"
	"orb-api/services"
)

type Controllers struct {
	User user.BaseHandler
}

func SetupControllers(service *services.Service) *Controllers {
	validator := config.NewValidator()

	return &Controllers{
		User: *user.NewBaseHandler(&service.User, validator),
	}
}
