package repositories

import (
	"gorm.io/gorm"
	"orb-api/repositories/user_role"
)

type Repository struct {
	DB       *gorm.DB
	UserRole UserRolesRepo.RUserRole
}

func SetupRepository(db *gorm.DB) Repository {
	var repo Repository
	repo = Repository{
		DB:       db,
		UserRole: UserRolesRepo.Setup(&repo),
	}

	// call exemple:
	//	repo.UserRole.ReadByUser(
	//		UserRolesRepo.IReadByUser{
	//			UserId: 10,
	//		})

	return repo
}
