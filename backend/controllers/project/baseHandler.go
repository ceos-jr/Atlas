package project

import (
	"orb-api/config"
	"orb-api/services/project"
)

type BaseHandler struct {
	Service   *project.Service
	Validator *config.RequestBodyValidator
}

func NewBaseHandler(
	service *project.Service,
	validator *config.RequestBodyValidator,
) *BaseHandler {
	return &BaseHandler{
		Service:   service,
		Validator: validator,
	}
}
