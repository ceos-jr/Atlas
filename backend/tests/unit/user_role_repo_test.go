package unit

import (
	"github.com/stretchr/testify/assert"
	"math/rand"
	"orb-api/config"
	"orb-api/repositories"
	"orb-api/repositories/seeds"
	iUserRole "orb-api/repositories/user_role/user_role_interface"
	"testing"
)

var (
	rawRepo, setupDBError = config.SetupDB()
	globalRepo            = repositories.SetupRepository(rawRepo.DB)
	users, usersSeedError = seeds.UserRandSeed(globalRepo.DB, 10)
	roles, rolesSeedError = seeds.RoleRandSeed(globalRepo.DB, 10)
)

func TestUserRoleCreate(test *testing.T) {
	assert := assert.New(test)

	randUserIndex := rand.Intn(len(*users))
	randRoleIndex := rand.Intn(len(*roles))

	result := globalRepo.UserRole.Create(
		iUserRole.ICreateUserRole{
			UserId: (*users)[randUserIndex].ID,
			RoleId: (*roles)[randRoleIndex].ID,
		},
	)

	assert.Nil(result)
}

//func TestUserRoleDelete(test *testing.T) {
//	assert := assert.New(test)
//	var arbitraryId uint = 0
//
//	assert.Nil(globalRepo.UserRole.Delete(
//		iUserRole.IDeleteUserRole{
//			UserRoleId: arbitraryId,
//		},
//	))
//}
