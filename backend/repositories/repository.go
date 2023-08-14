package repository

import (
	"gorm.io/gorm"
	"orb-api/repositories/user_role"
)

type Repository struct {
	DB       *gorm.DB
	UserRole userrole.Repository
}

func SetupRepository(connection *gorm.DB) *Repository {
	return &Repository{
		DB:       connection,
		UserRole: userrole.NewUserRoleRepository(connection),
	}
}
