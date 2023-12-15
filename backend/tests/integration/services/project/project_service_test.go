package projectservicetest

import (
	"orb-api/models"
	"orb-api/services/project"
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

func (suite *TestSuit) TestUpdateProject () {
	newName := "projectupdated"
	newSector := uint(3)
	newAdmID := uint(5)
	
	UpdatedProject, updateErr := suite.ProjectService.UpdateProject(project.Update{
		ID: suite.MockProjects[1].ID,
		Name: &newName,
		Sector: &newSector,
		AdmID: &newAdmID,
	})
	
	suite.Nil(updateErr, "Update error must be nil")
	
	suite.Equal(UpdatedProject.ID, suite.MockProjects[1].ID)
	suite.Equal(UpdatedProject.Name, newName)
	suite.Equal(UpdatedProject.Sector, newSector)
	suite.Equal(UpdatedProject.AdmID, newAdmID)
}

func (suite *TestSuit) TestUpdateProjectErr() {
	id := uint(1)
	min := "aa"
	max := "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
	newName := "projectupdated"
	
	_, err := suite.ProjectService.UpdateProject(project.Update{
		ID: id,
		Name: &newName,
	})
	suite.Equal("Invalid Project ID", err.Error(), "expected to have an error")

	_, err2 := suite.ProjectService.UpdateProject(project.Update{
		ID: suite.MockProjects[2].ID,
		Name: &min,
	})
	suite.Equal("Invalid project name", err2.Error(), "expected to have an error")

	_, err3 := suite.ProjectService.UpdateProject(project.Update{
		ID: suite.MockProjects[3].ID,
		Name: &max,
	})
	suite.Equal("Invalid project name", err3.Error(), "expected to have an error")
}

func TestProjectService(test *testing.T) {
	suite.Run(test, new(TestSuit))
}
