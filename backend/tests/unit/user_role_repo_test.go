package unit

import (
	"math/rand"
	"orb-api/config"
	"orb-api/repositories/seeds"
	"orb-api/repositories/user_role"
	"testing"
	"github.com/stretchr/testify/assert"
)

var (
	repository, setupDBError = config.SetupDB()
	users, usersSeedError = seeds.UserRandSeed(repository.DB, 10)
	roles, rolesSeedError = seeds.RoleRandSeed(repository.DB, 10)
)

func TestUserRoleCreate(test *testing.T) {
	assert := assert.New(test)

	randUserIndex := rand.Intn(len(*users))
	randRoleIndex := rand.Intn(len(*roles))

	result := repository.UserRole.Create(
		userrole.ICreateUserRole{
			UserId: (*users)[randUserIndex].ID,
			RoleId: (*roles)[randRoleIndex].ID,
		},
	)

	assert.Nil(result)
}
