package sector

import (
	"orb-api/config"
	"orb-api/services/sector"
)

type BaseHandler struct {
	Service   *sector.Service
	Validator *config.RequestBodyValidator
}

func NewBaseHandler(
	service *sector.Service,
	validator *config.RequestBodyValidator,
) *BaseHandler {
	return &BaseHandler{
		Service:   service,
		Validator: validator,
	}
}
