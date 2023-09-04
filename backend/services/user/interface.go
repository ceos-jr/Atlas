package userservice

import (
	"orb-api/models"
	"orb-api/repositories/user"
)

const (
	emailMaxLen    = 128
	emailMinLen    = 3
	nameMaxLen     = 128
	nameMinLen     = 5
	passwordMinLen = 8
)

type (
	UserService struct {
		UserRepo user.Repository
	}

	ICreateUser struct {
		Name     string
		Email    string
		Password string
	}

	Interface interface {
		CreateNewUser(ICreateUser) (*models.User, error)
	}
)
