package userservicerepotest

import (
	"orb-api/models"

	"github.com/stretchr/testify/suite"
)

type UserServiceRepoTestSuit struct{
	suite.Suite
	MockUsers []models.User
	MockTasks []models.Task
}