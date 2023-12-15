package user_role

import (
	"orb-api/config"
	"orb-api/services/user_role"
)

type BaseHandler struct {
	Service   *user_role.Service
	Validator *config.RequestBodyValidator
}

func NewBaseHandler(
	service *user_role.Service,
	validator *config.RequestBodyValidator,
) *BaseHandler {
	return &BaseHandler{
		Service:   service,
		Validator: validator,
	}
}
