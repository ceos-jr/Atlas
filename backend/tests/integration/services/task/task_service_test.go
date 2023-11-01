package taskservicetest

import (
	"orb-api/services/task"
	"testing"

	"github.com/stretchr/testify/suite"
)

func (suite TestSuit) TestConcludedTask() {
	UpdateStatus, UpdateErr := suite.Service.ConcludedTask(suite.MockTasks[0].ID)

	suite.Nil(UpdateErr, "Update error must be Nil")

	suite.Equal(UpdateStatus.Status, 1)
}

