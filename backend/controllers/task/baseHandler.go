package task

import (
	"orb-api/config"
	"orb-api/services/task"
)

type BaseHandler struct {
	Service   *task.Service
	Validator *config.RequestBodyValidator
}

func NewBaseHandler(
	service *task.Service,
	validator *config.RequestBodyValidator,
) *BaseHandler {
	return &BaseHandler{
		Service:   service,
		Validator: validator,
	}
}
