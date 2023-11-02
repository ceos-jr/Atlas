package taskservicetest

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

func (suite *TestSuit) TestConcludedTask() {
	NewStatus := uint(1)

	UpdateStatus, UpdateErr := suite.TaskService.MarkTaskAsCompleted(suite.MockTasks[0].ID)
  
	suite.Nil(UpdateErr, "Update error must be nil")

	suite.Equal(NewStatus, UpdateStatus.Status)
}

func TestTaskRepository(t *testing.T) {
	suite.Run(t, new(TestSuit))
}
