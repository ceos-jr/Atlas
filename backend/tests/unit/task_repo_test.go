package unit

import (
	"fmt"
	"orb-api/config"
	"orb-api/models"
	"orb-api/repositories"
	"orb-api/repositories/task"
	"orb-api/repositories/user"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
)

type TaskRepoTestSuite struct {
  suite.Suite 
  Repo *repository.Repository
  MockUsers []models.User
  MockTasks []models.Task
}

// Executed before all tests
func (suite *TaskRepoTestSuite) SetupSuite() {
  repo, setupError := config.SetupDB() 
  
  if setupError != nil {
    panic(setupError)
  }

  suite.Repo = repo
  suite.SetupMocks()
}

// Executed after all tests 
func (suite *TaskRepoTestSuite) TearDownSuite() {
  for index := range suite.MockUsers {
    _, deleteErr := suite.Repo.User.Delete(user.IDelete{
      ID: suite.MockUsers[index].ID,
    })

    if deleteErr != nil {
      panic(deleteErr)
    }
  }

  for index := range suite.MockTasks {
    _, deleteErr := suite.Repo.Task.Delete(task.IDelete{
      ID: suite.MockTasks[index].ID,
    })

    if deleteErr != nil {
      panic(deleteErr)
    }
  }
}

func (suite *TaskRepoTestSuite) SetupMocks() {
  var users = make([]models.User, 2) 

  for index := range users {
    user, createErr := suite.Repo.User.Create(user.ICreate{
      Name: fmt.Sprintf("Gabrigas %v", index + 1),
      Email: fmt.Sprintf("example0%v@example.com", index + 1),
      Password: "gabrigas123",
      Status: 2,
    })
    
    if createErr != nil {
      panic(createErr)
    }
    
    users[index] = *user
  }
  
  suite.MockUsers = users
}

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
  suite.Equal(
    time.Date(2077, 4, 12, 12, 0, 0, 0, time.UTC), task.Deadline,
    "Deadlines do not match",
  )
  
  tasks[0] = *task

  suite.MockTasks = tasks
}

func TestTaskRepository(test *testing.T) {
  suite.Run(test, new(TaskRepoTestSuite))
}
