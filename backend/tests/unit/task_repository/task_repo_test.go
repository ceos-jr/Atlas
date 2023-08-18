package taskrepotest

import (
	"orb-api/models"
	"orb-api/repositories/task"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
)

func (suite *TaskRepoTestSuite) TestCreateTask() {
  var tasks = make([]models.Task, 1)
  
  task, createErr := suite.Repo.Task.Create(task.ICreate{
    Description: "This is a test",
    CreatedBy: suite.MockUsers[0].ID,
    AssignedTo: suite.MockUsers[1].ID,
    Status: 2,
    Deadline: time.Date(2077, 4, 12, 12, 0, 0, 0, time.UTC),
  })  

  suite.Nil(createErr, "Create error must be nil")
  suite.Equal("This is a test", task.Description, "Description do not match")
  suite.Equal(suite.MockUsers[0].ID, task.CreatedBy, "CreatedBy do not match")
  suite.Equal(suite.MockUsers[1].ID, task.AssignedTo, "AssignedTo do not match")
  suite.Equal(uint(2), task.Status, "Status do not match")
  suite.Equal(time.Date(2077, 4, 12, 12, 0, 0, 0, time.UTC), task.Deadline,
    "Deadlines do not match",
  )
  
  tasks[0] = *task

  suite.MockTasks = tasks
}

func (suite *TaskRepoTestSuite) TestCreateTaskErr() {
  _, createErr := suite.Repo.Task.Create(task.ICreate{
    Description: "This is a test",
    CreatedBy: 999,
    AssignedTo: suite.MockUsers[1].ID,
    Status: 2,
    Deadline: time.Date(2077, 4, 12, 12, 0, 0, 0, time.UTC),
  })

  suite.Equal("record not found", createErr.Error(), 
    "Createby invalid, it should return an error",
  )
 
  _, createErr = suite.Repo.Task.Create(task.ICreate{
    Description: "This is a test",
    CreatedBy: suite.MockUsers[0].ID,
    AssignedTo: 999,
    Status: 2,
    Deadline: time.Date(2077, 4, 12, 12, 0, 0, 0, time.UTC),
  })

  suite.Equal("record not found", createErr.Error(), 
    "AssignedTo invalid, it should return an error",
  )
  
  _, createErr = suite.Repo.Task.Create(task.ICreate{
    Description: "This is a test",
    CreatedBy: suite.MockUsers[0].ID,
    AssignedTo: suite.MockUsers[1].ID,
    Status: 77,
    Deadline: time.Date(2077, 4, 12, 12, 0, 0, 0, time.UTC),
  })

  suite.Equal("Invalid task status", createErr.Error(), 
    "status invalid, it should return an error",
  )

  _, createErr = suite.Repo.Task.Create(task.ICreate{
    Description: "This is a test",
    CreatedBy: suite.MockUsers[0].ID,
    AssignedTo: suite.MockUsers[1].ID,
    Status: 2,
    Deadline: time.Date(2004, 4, 12, 12, 0, 0, 0, time.UTC),
  })

  suite.Equal("Invalid deadline", createErr.Error(), 
    "Deadline invalid, it should return an error",
  )
}

func (suite *TaskRepoTestSuite) TestReadAllTasks() {
  tasks, readErr := suite.Repo.Task.ReadAll()  

  suite.Nil(readErr, "Read error must be nil")
  suite.Equal(1, len(tasks), "Expected to have one task")
  suite.Equal(suite.MockTasks[0].ID, tasks[0].ID, 
    "Expected to have the same id",
  ) 
}

func (suite *TaskRepoTestSuite) TestDeleteTask() {
  newtask, _ := suite.Repo.Task.Create(task.ICreate{
    Description: "This is another test",
    CreatedBy: suite.MockUsers[0].ID,
    AssignedTo: suite.MockUsers[1].ID,
    Status: 2,
    Deadline: time.Date(2050, 4, 12, 12, 0, 0, 0, time.UTC),
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
