package userrole

import (
	"github.com/stretchr/testify/assert"
	"orb-api/config"
	"orb-api/repositories/role"
	"orb-api/repositories/seeds"
	"orb-api/repositories/user"
	"orb-api/repositories/user_role"
	"testing"
)

var (
	repository, setupDBError = config.SetupDB()
)

func TestUserRoleCreate(test *testing.T) {
	assert := assert.New(test)
	testSize := 10

	_, createUserError := seeds.UserRandSeed(repository, testSize)
	_, createRoleError := seeds.RoleRandSeed(repository, testSize)

	assert.Nil(createUserError)
	assert.Nil(createRoleError)

	readUsers, readUsersError := repository.User.ReadAll(user.IReadAll{Limit: &testSize})
	readRoles, readRolesError := repository.Role.ReadAll(role.IReadAll{Limit: &testSize})

	assert.Nil(readUsersError)
	assert.Nil(readRolesError)

	userRoles := make([]userrole.ICreate, testSize)

	for i := range userRoles {
		readUser := readUsers[i]
		readRole := readRoles[i]

		userRoles[i] = userrole.ICreate{
			RoleID: readRole.ID,
			UserID: readUser.ID,
		}

		result := repository.UserRole.Create(userRoles[i])

		assert.Nil(result)
	}
}
