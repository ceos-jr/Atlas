package userservice 

import (
  "orb-api/repositories/user"
  "orb-api/repositories"
)

type UserService struct {
  userRepo user.Repository; 
}

func SetupUserService(repo repository.Repository) *UserService {
  return &UserService{
    userRepo: repo.User,
  }
} 
