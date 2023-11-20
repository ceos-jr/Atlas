package user_roleservicetest

import (
	"fmt"
	"orb-api/config"
	"orb-api/models"
	rolerepo "orb-api/repositories/role"
	userrepo "orb-api/repositories/user"
	user_rolerepo "orb-api/repositories/user_role"

	user_roleservice "orb-api/services/user_role"

	"github.com/stretchr/testify/suite"
)

type TestSuite struct {
	suite.Suite
	Service   *user_roleservice.Service
	MockUsers []models.User
	MockRoles []models.Role
	MockUserRoles []models.UserRole
}

// SetupSuite Executed before all tests
func (suite *TestSuite) SetupSuite(){
	repositories, setupError := config.SetupDB("../../.env")

	if setupError != nil {
		panic(setupError)
	}

	suite.Service = user_roleservice.SetupService(&repositories.User, &repositories.Role, &repositories.UserRole)
	suite.MockUserRoles = make([]models.UserRole, 1)
	suite.MockUsers = make([]models.User, 2)
	suite.MockRoles = make([]models.Role, 2)
	suite.SetupMocks()
}

// setting up the mock UserRole
func (suite *TestSuite) SetupMocks() {

	//criar os mocks de user
	for i := 0; i < 2; i++ {
		NewUser, createErr := suite.Service.UserRepo.Create(userrepo.ICreate{
			Name:     fmt.Sprintf("User 0%v", i+1),
			Email:    fmt.Sprintf("example0%v@example.com", i+1),
			Password: "gabrigas123",
			Status:   2,
		})

		if createErr != nil {
			panic(createErr)
		}

		suite.MockUsers[i] = *NewUser
	}
	//Criar os mocks de role
	for i := 0; i < 2; i++ {
		newRole, createErr := suite.Service.RoleRepo.Create(rolerepo.ICreate{
			Name:     fmt.Sprintf("Role 0%v", i+1),
			Description:     fmt.Sprintf("description 0%v", i+1),
		})

		if createErr != nil {
			panic(createErr)
		}

		suite.MockRoles[i] = *newRole
	}

	//Criação do mock UserRole
	newUserRole, errs := suite.Service.UserRoleRepo.Create(user_rolerepo.ICreate{
		RoleID: suite.MockRoles[0].ID,
		UserID: suite.MockUsers[0].ID,
	})
	if errs != nil{
		panic(errs)
	}

	suite.MockUserRoles[0] = *newUserRole
}

// TearDownSuite Executed after all tests
func (suite *TestSuite) TearDownSuite() {
	
	_, deleteErr := suite.Service.UserRepo.Delete(userrepo.IDelete{
		ID: suite.MockUsers[0].ID,
	})
	
	if deleteErr != nil {
		panic(deleteErr)
	}
	
	_, deleteErr = suite.Service.RoleRepo.Delete(rolerepo.IDelete{
		RoleID: suite.MockRoles[0].ID,
	})
	
	if deleteErr != nil {
		panic(deleteErr)
	}
	
	_, deleterrs := suite.Service.UserRoleRepo.Delete(user_rolerepo.IDelete{
		UserRoleID: suite.MockUserRoles[0].ID,
	})
	
	if deleterrs != nil{
		panic(deleterrs)
	}
}