package projectservicetest

import (
	"orb-api/models"
	"testing"

	"github.com/stretchr/testify/suite"
)

func (suite *TestSuit) TestSortByDeadline () {
	var taskprojects = make([]models.TasksProject, 2)

	t_project1, assignErr := suite.ProjectService.AssignTask(suite.MockProjects[0].ID, suite.MockTasks[0].ID)

	suite.Nil(assignErr, "Assign error must be nil")

	t_project2, assignErr := suite.ProjectService.AssignTask(suite.MockProjects[0].ID, suite.MockTasks[1].ID)

	suite.Nil(assignErr, "Assign error must be nil")

	taskprojects[0] = *t_project1
	taskprojects[1] = *t_project2

	suite.MockTasksProjects = taskprojects

	SortedTasks, SortErr := suite.ProjectService.SortByDeadline(suite.MockProjects[0].ID)

	suite.Nil(SortErr, "Sorting error must be nil")

	suite.Equal(SortedTasks[0].Description, suite.MockTasks[1].Description, "First task Description does not match")
	suite.Equal(SortedTasks[0].AssignedTo, suite.MockTasks[1].AssignedTo, "First task AssignedTo does not match")
	suite.Equal(SortedTasks[0].CreatedBy, suite.MockTasks[1].CreatedBy, "First task CreatedBy does not match")
	suite.Equal(SortedTasks[0].Status, suite.MockTasks[1].Status, "First task Status does not match")
	suite.Equal(SortedTasks[0].Deadline, suite.MockTasks[1].Deadline, "First task Deadline does not match")

	suite.Equal(SortedTasks[1].Description, suite.MockTasks[0].Description, "Second task Description does not match")
	suite.Equal(SortedTasks[1].AssignedTo, suite.MockTasks[0].AssignedTo, "Second task AssignedTo does not match")
	suite.Equal(SortedTasks[1].CreatedBy, suite.MockTasks[0].CreatedBy, "Second task CreatedBy does not match")
	suite.Equal(SortedTasks[1].Status, suite.MockTasks[0].Status, "Second task Status does not match")
	suite.Equal(SortedTasks[1].Deadline, suite.MockTasks[0].Deadline, "Second task Deadline does not match")
}

func TestProjectService(test *testing.T) {
	suite.Run(test, new(TestSuit))
}
