package repositories

import (
	"gorm.io/gorm"
	. "orb-api/repositories/globalInterface"
)

func SetupRepository(db *gorm.DB) Repository {
	var repo Repository
	repo = Repository{
		DB:       db,
		UserRole: SetupUserRole(&repo),
	}

	// call exemple:
	//	repo.UserRole.ReadByUser(
	//		userRole.IReadByUser{
	//			UserId: 10,
	//		})

	return repo
}
