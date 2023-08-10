package globalInterface

import (
	"gorm.io/gorm"
	userRole "orb-api/repositories/user_role"
)

type (
	Repository struct {
		DB       *gorm.DB
		UserRole userRole.RUserRole
	}
)

func SetupUserRole(repo *Repository) userRole.RUserRole {
	return userRole.RUserRole{
		Repo: func() *gorm.DB {
			return repo.DB
		},
	}
}
