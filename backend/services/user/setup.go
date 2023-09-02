package userservice

import (
	"orb-api/repositories"
	"orb-api/repositories/user"
)

type UserService struct {
	userRepo user.Repository
}

func SetupUserService(repo repository.Repository) *UserService {
	return &UserService{
		userRepo: repo.User,
	}
}
