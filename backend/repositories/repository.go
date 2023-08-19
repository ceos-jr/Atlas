package repository

import (
	"gorm.io/gorm"
	"orb-api/repositories/role"
	"orb-api/repositories/task"
	"orb-api/repositories/user"
	"orb-api/repositories/user_role"
)

type Repository struct {
	DB       *gorm.DB
	UserRole userrole.Repository
	Role     role.Repository
	User     user.Repository
	Task     task.Repository
}

func SetupRepository(connection *gorm.DB) *Repository {
	return &Repository{
		DB:       connection,
		UserRole: userrole.NewUserRoleRepository(connection),
		Role:     role.NewRoleRepository(connection),
		User:     user.NewUserRepository(connection),
		Task:     task.NewTaskRepository(connection),
	}
}
