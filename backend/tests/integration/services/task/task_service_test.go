package taskservicetest

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

func (suite *TestSuit) TestConcludedTask() {
	NewStatus := uint(1)

	UpdateStatus, UpdateErr := suite.TaskService.MarkTaskAsCompleted(suite.MockTasks[0].ID)

	suite.Nil(UpdateErr, "Update error must be nil")

	suite.Equal(NewStatus, UpdateStatus.Status)
}

func (suite *TestSuit) TestAssignTask() {

	AssignedTask, AssignErr := suite.TaskService.AssignTask(suite.MockTasks[0].ID, suite.MockUsers[1].ID)

	suite.Nil(AssignErr, "Assing error must be Nil")

	suite.Equal(AssignedTask.AssignedTo, suite.MockUsers[1].ID)
}


func (suite *TestSuit) TestOrganizeTasks() {

	OrganizedTasks, OrganizeErr := suite.TaskService.OrganizeTasks(suite.MockUsers[2].ID)

	suite.Nil(OrganizeErr, "Organize error must be Nil")
	
	suite.Equal((*OrganizedTasks)[1].Description, suite.MockTasks[1].Description)
	suite.Equal((*OrganizedTasks)[0].Description, suite.MockTasks[2].Description)
}

func (suite *TestSuit) TestOrganizeTasksErr() {
	id := uint(524)

	_, err := suite.TaskService.OrganizeTasks(id)
	suite.Equal("invalid user id", err.Error(), "Expected to have an error")
}

func TestTaskRepository(t *testing.T) {
	suite.Run(t, new(TestSuit))
}
