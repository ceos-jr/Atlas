package project

import (
	"github.com/gofiber/fiber/v2"
)

type CreateUserRequestBodyId struct {
	Id     uint `json:"id" validate:"required"`
}

