package repository

import (
	"orb-api/repositories/message"
	project "orb-api/repositories/project"
	"orb-api/repositories/relation"
	"orb-api/repositories/role"
	"orb-api/repositories/sector"
	"orb-api/repositories/task"
	"orb-api/repositories/taskproject"
	"orb-api/repositories/user"
	userrole "orb-api/repositories/user_role"
	"orb-api/repositories/userproject"

	"gorm.io/gorm"
)

type Repository struct {
	DB          *gorm.DB
	UserRole    userrole.Repository
	Role        role.Repository
	User        user.Repository
	Task        task.Repository
	Message     message.Repository
	Relation    relation.Repository
	Project     project.Repository
	UserProject userproject.Repository
	TaskProject taskproject.Repository
	Sector      sector.Repository
}

func SetupRepository(connection *gorm.DB) *Repository {
	return &Repository{
		DB:          connection,
		UserRole:    userrole.NewUserRoleRepository(connection),
		Role:        role.NewRoleRepository(connection),
		User:        user.NewUserRepository(connection),
		Task:        task.NewTaskRepository(connection),
		Message:     message.NewMessageRepository(connection),
		Relation:    relation.NewRepository(connection),
		Project:     project.NewProjectRepository(connection),
		UserProject: userproject.NewUserProjectRepository(connection),
		TaskProject: taskproject.NewTaskProjectRepository(connection),
		Sector:      sector.NewSectorRepository(connection),
	}
}
