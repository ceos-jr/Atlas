package taskrepotest

import (
	"orb-api/repositories/task"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
)

func (suite *TaskRepoTestSuite) TestCreateTask() {
	task, createErr := suite.Repo.Task.Create(task.ICreate{
		Description: "This is a test",
		CreatedBy:   suite.MockUsers[0].ID,
		AssignedTo:  suite.MockUsers[1].ID,
		Status:      2,
		Deadline:    time.Date(2077, 4, 12, 12, 0, 0, 0, time.UTC),
	})

	suite.Nil(createErr, "Create error must be nil")
	suite.Equal("This is a test", task.Description, "Description do not match")
	suite.Equal(suite.MockUsers[0].ID, task.CreatedBy, "CreatedBy do not match")
	suite.Equal(suite.MockUsers[1].ID, task.AssignedTo, "AssignedTo do not match")
	suite.Equal(uint(2), task.Status, "Status do not match")
	suite.Equal(time.Date(2077, 4, 12, 12, 0, 0, 0, time.UTC), task.Deadline,
		"Deadlines do not match",
	)

	suite.MockTasks[1] = *task
}

func (suite *TaskRepoTestSuite) TestCreateTaskErr() {
	_, createErr := suite.Repo.Task.Create(task.ICreate{
		Description: "This is a test",
		CreatedBy:   999,
		AssignedTo:  suite.MockUsers[1].ID,
		Status:      2,
		Deadline:    time.Date(2077, 4, 12, 12, 0, 0, 0, time.UTC),
	})

	suite.Equal("Invalid user passed to createBy", createErr.Error(),
		"Createby invalid, it should return an error",
	)

	_, createErr = suite.Repo.Task.Create(task.ICreate{
		Description: "This is a test",
		CreatedBy:   suite.MockUsers[0].ID,
		AssignedTo:  999,
		Status:      2,
		Deadline:    time.Date(2077, 4, 12, 12, 0, 0, 0, time.UTC),
	})

	suite.Equal("Invalid user passed to assignedTo", createErr.Error(),
		"AssignedTo invalid, it should return an error",
	)

	_, createErr = suite.Repo.Task.Create(task.ICreate{
		Description: "This is a test",
		CreatedBy:   suite.MockUsers[0].ID,
		AssignedTo:  suite.MockUsers[1].ID,
		Status:      77,
		Deadline:    time.Date(2077, 4, 12, 12, 0, 0, 0, time.UTC),
	})

	suite.Equal("Invalid task status", createErr.Error(),
		"status invalid, it should return an error",
	)

	_, createErr = suite.Repo.Task.Create(task.ICreate{
		Description: "This is a test",
		CreatedBy:   suite.MockUsers[0].ID,
		AssignedTo:  suite.MockUsers[1].ID,
		Status:      2,
		Deadline:    time.Date(2004, 4, 12, 12, 0, 0, 0, time.UTC),
	})

	suite.Equal("Invalid deadline", createErr.Error(),
		"Deadline invalid, it should return an error",
	)
}

func (suite *TaskRepoTestSuite) TestReadAllTasks() {
	tasks, readError := suite.Repo.Task.ReadAll()

	suite.Nil(readError, "Read error must be nil")
	suite.Equal(2, len(tasks), "Expected to have two tasks")
	suite.Equal(suite.MockTasks[0].ID, tasks[0].ID,
		"Expected to have the same id",
	)
}

func (suite *TaskRepoTestSuite) TestReadTaskByID() {
	tasks, readError := suite.Repo.Task.ReadBy(task.IReadBy{
		ID: &suite.MockTasks[0].ID,
	})

	suite.Nil(readError, "Read error should be empty")
	suite.Equal(suite.MockTasks[0].ID, tasks[0].ID, "The ids must match")
}

func (suite *TaskRepoTestSuite) TestReadTaskByAssigned() {
	tasks, readError := suite.Repo.Task.ReadBy(task.IReadBy{
		AssignedTo: &suite.MockTasks[0].AssignedTo,
	})

	suite.Nil(readError, "Read error should be empty")
	suite.Equal(suite.MockTasks[0].AssignedTo, tasks[0].AssignedTo,
		"The user assigned to the task must match",
	)
}

func (suite *TaskRepoTestSuite) TestReadTaskByCreator() {
	tasks, readError := suite.Repo.Task.ReadBy(task.IReadBy{
		CreatedBy: &suite.MockTasks[0].CreatedBy,
	})

	suite.Nil(readError, "Read error should be empty")
	suite.Equal(suite.MockTasks[0].CreatedBy, tasks[0].CreatedBy,
		"The creator of the task must match",
	)
}

func (suite *TaskRepoTestSuite) TestReadTaskByStatus() {
	tasks, readError := suite.Repo.Task.ReadBy(task.IReadBy{
		Status: &suite.MockTasks[0].Status,
	})

	suite.Nil(readError, "Read error should be empty")
	suite.Equal(suite.MockTasks[0].Status, tasks[0].Status,
		"The status must match",
	)
}

func (suite *TaskRepoTestSuite) TestReadTaskByTimeRange() {
	timeRange := time.Date(2100, 12, 4, 12, 0, 0, 0, time.UTC)

	tasks, readError := suite.Repo.Task.ReadBy(task.IReadBy{
		TimeRange: &timeRange,
	})

	suite.Nil(readError, "Read error should be empty")
	suite.Equal(true, tasks[0].Deadline.Before(timeRange),
		"Deadline must be in the time range",
	)
}

func (suite *TaskRepoTestSuite) TestReadByErr() {
	_, readError := suite.Repo.Task.ReadBy(task.IReadBy{})

	suite.Equal("No fields to read", readError.Error(),
		"Empty fields it should return an error",
	)

	var invalidStatus uint = 777

	_, readError = suite.Repo.Task.ReadBy(task.IReadBy{
		Status: &invalidStatus,
	})

	suite.Equal("Invalid task status", readError.Error(),
		"Invalid status it should return an error",
	)
}

func (suite *TaskRepoTestSuite) TestUpdateTask() {
	description := "Updated task"
	status := uint(3)
	deadline := time.Date(2099, 12, 4, 12, 0, 0, 0, time.UTC)

	updatedTask, updateError := suite.Repo.Task.Update(task.IUpdate{
		ID:          suite.MockTasks[1].ID,
		Description: &description,
		CreatedBy:   &suite.MockUsers[1].ID,
		AssignedTo:  &suite.MockUsers[0].ID,
		Status:      &status,
		Deadline:    &deadline,
	})

	suite.Nil(updateError, "Update error must be nil")
	suite.Equal(description, updatedTask.Description,
		"Descriptions do not match",
	)
	suite.Equal(suite.MockUsers[1].ID, updatedTask.CreatedBy,
		"CreatedBy do not match",
	)
	suite.Equal(suite.MockUsers[0].ID, updatedTask.AssignedTo,
		"AssignedTo do not match",
	)
	suite.Equal(status, updatedTask.Status,
		"Status do not match",
	)
	suite.Equal(deadline, updatedTask.Deadline,
		"deadlines do not match",
	)
}

func (suite *TaskRepoTestSuite) TestUpdateTaskErr() {
	invalidStatus := uint(77)
	invalidDeadline := time.Date(2012, 12, 4, 12, 0, 0, 0, time.UTC)
	invalidUserID := uint(777)

	_, updateError := suite.Repo.Task.Update(task.IUpdate{
		ID: suite.MockTasks[1].ID,
	})

	suite.Equal("No fields to update", updateError.Error(),
		"Empty fields it should return an error",
	)

	_, updateError = suite.Repo.Task.Update(task.IUpdate{
		ID:     suite.MockTasks[1].ID,
		Status: &invalidStatus,
	})

	suite.Equal("Invalid task status", updateError.Error(),
		"Invalid task status it should return an error",
	)

	_, updateError = suite.Repo.Task.Update(task.IUpdate{
		ID:       suite.MockTasks[1].ID,
		Deadline: &invalidDeadline,
	})

	suite.Equal("Invalid deadline", updateError.Error(),
		"Invalid task deadline it should return an error",
	)

	_, updateError = suite.Repo.Task.Update(task.IUpdate{
		ID:        suite.MockTasks[1].ID,
		CreatedBy: &invalidUserID,
	})

	suite.Equal("Invalid user passed to createBy", updateError.Error(),
		"Invalid createby it should return an error",
	)

	_, updateError = suite.Repo.Task.Update(task.IUpdate{
		ID:         suite.MockTasks[1].ID,
		AssignedTo: &invalidUserID,
	})

	suite.Equal("Invalid user passed to assignedTo", updateError.Error(),
		"Invalid assignedto it should return an error",
	)
}

func (suite *TaskRepoTestSuite) TestDeleteTask() {
	newtask, _ := suite.Repo.Task.Create(task.ICreate{
		Description: "This is another test",
		CreatedBy:   suite.MockUsers[0].ID,
		AssignedTo:  suite.MockUsers[1].ID,
		Status:      2,
		Deadline:    time.Date(2050, 4, 12, 12, 0, 0, 0, time.UTC),
	})

	deletedtask, deleteErr := suite.Repo.Task.Delete(task.IDelete{
		ID: newtask.ID,
	})

	suite.Nil(deleteErr, "Delete error must be nil")
	suite.Equal(newtask.ID, deletedtask.ID, "Expected to have the same id")
}

func TestTaskRepository(test *testing.T) {
	suite.Run(test, new(TaskRepoTestSuite))
}
