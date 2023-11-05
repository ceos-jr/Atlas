package taskservicetest

import (

	"github.com/stretchr/testify/suite"
	"testing"
)

func (suite *TestSuit) TestAssignTask() {

	AssignedTask, AssignErr := suite.TaskService.AssignTask(suite.MockTasks[0].ID, suite.MockUsers[2].ID)

	suite.Nil(AssignErr, "Assing error must be Nil")

	suite.Equal(AssignedTask.AssignedTo, suite.MockUsers[2].ID)
}

func TestTaskService(test *testing.T) {
	suite.Run(test, new(TestSuit))
}