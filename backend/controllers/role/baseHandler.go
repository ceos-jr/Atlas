package role

import (
	"orb-api/config"
	"orb-api/services/role"
)

type BaseHandler struct {
	Service   *role.Service
	Validator *config.RequestBodyValidator
}

func NewBaseHandler(
	service *role.Service,
	validator *config.RequestBodyValidator,
) *BaseHandler {
	return &BaseHandler{
		Service:   service,
		Validator: validator,
	}
}
