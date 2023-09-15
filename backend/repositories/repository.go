package repository

import (
	"orb-api/repositories/message"
	"orb-api/repositories/relation"
	"orb-api/repositories/role"
	"orb-api/repositories/task"
	"orb-api/repositories/user"
	userrole "orb-api/repositories/user_role"

	"gorm.io/gorm"
)

type Repository struct {
	DB       *gorm.DB
	UserRole userrole.Repository
	Role     role.Repository
	User     user.Repository
	Task     task.Repository
	Message  message.Repository
	Relation relation.Repository
}

func SetupRepository(connection *gorm.DB) *Repository {
	return &Repository{
		DB:       connection,
		UserRole: userrole.NewUserRoleRepository(connection),
		Role:     role.NewRoleRepository(connection),
		User:     user.NewUserRepository(connection),
		Task:     task.NewTaskRepository(connection),
		Message:  message.NewMessageRepository(connection),
		Relation: relation.NewRepository(connection),
	}
}
