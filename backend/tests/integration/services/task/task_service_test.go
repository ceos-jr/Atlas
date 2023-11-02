package taskservicetest

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

func (suite TestSuit) TestConcludedTask() {
	NewStatus := uint(1)

	UpdateStatus, UpdateErr := suite.TaskService.ConcludedTask(suite.MockTasks[0].ID)

	suite.Nil(UpdateErr, "Update error must be nil")

	suite.Equal(UpdateStatus.Status, NewStatus)
}

func TestTaskRepository(t *testing.T) {
	suite.Run(t, new(TestSuit))
}