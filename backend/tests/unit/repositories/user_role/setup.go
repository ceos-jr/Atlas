package userolerepotest

import (
	"orb-api/config"
	"orb-api/models"
	repository "orb-api/repositories"
	"orb-api/repositories/role"
	"orb-api/repositories/user"
	userrole "orb-api/repositories/user_role"

	"github.com/bxcodec/faker/v3"
	"github.com/stretchr/testify/suite"
)

const (
	MockArraySize = 2
)

type UserRoleRepoTestSuite struct {
	suite.Suite
	Repo         *repository.Repository
	MockUser     []models.User
	MockRole     []models.Role
	MockUserRole []models.UserRole
}

func (s *UserRoleRepoTestSuite) SetupSuite() {
	repo, setupErr := config.SetupDB("../../.env")

	if setupErr != nil {
		panic(setupErr)
	}

	s.Repo = repo
	s.MockUser = make([]models.User, 3)
	s.MockRole = make([]models.Role, MockArraySize)
	s.MockUserRole = make([]models.UserRole, MockArraySize)
	s.SetupMocks()
}

func (s *UserRoleRepoTestSuite) TearDownSuite() {
	for index := range s.MockUser {
		_, deleteErr := s.Repo.User.Delete(user.IDelete{
			ID: s.MockUser[index].ID,
		})

		if deleteErr != nil {
			panic(deleteErr)
		}
	}
	for index := range s.MockRole {
		_, deleteErr := s.Repo.Role.Delete(role.IDelete{
			RoleID: s.MockRole[index].ID,
		})

		if deleteErr != nil {
			panic(deleteErr)
		}
	}

	for index := range s.MockUserRole {
		_, deleteErr := s.Repo.UserRole.Delete(userrole.IDelete{
			UserRoleID: s.MockUserRole[index].ID,
		})

		if deleteErr != nil {
			panic(deleteErr)
		}

	}
}

func (s *UserRoleRepoTestSuite) SetupMocks() {
	for index := 0; index < 3; index++ {
		createdUser, createError := s.Repo.User.Create(user.ICreate{
			Name:     faker.Name(),
			Email:    faker.Email(),
			Status:   models.UStatusActive,
			Password: faker.Password(),
		})

		if createError != nil {
			panic(createError)
		}

		s.MockUser[index] = *createdUser
	}

	for index := 0; index < 2; index++ {
		createdRole, createError := s.Repo.Role.Create(role.ICreate{
			Name:        faker.Name(),
			Description: faker.Sentence(),
		})

		if createError != nil {
			panic(createError)
		}

		s.MockRole[index] = *createdRole
	}

	userRole, createError := s.Repo.UserRole.Create(userrole.ICreate{
		RoleID: s.MockRole[0].ID,
		UserID: s.MockUser[0].ID,
	})

	if createError != nil {
		panic(createError)
	}

	s.MockUserRole[0] = *userRole

}
