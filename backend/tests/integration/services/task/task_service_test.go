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

func (suite *TestSuit) TestAssignTask() {

	AssignedTask, AssignErr := suite.TaskService.AssignTask(suite.MockTasks[0].ID, suite.MockUsers[2].ID)

	suite.Nil(AssignErr, "Assing error must be Nil")

	suite.Equal(AssignedTask.AssignedTo, suite.MockUsers[2].ID)
}

func TestTaskService(test *testing.T) {
	suite.Run(test, new(TestSuit))
}
