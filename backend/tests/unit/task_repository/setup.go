package taskrepotest 

import (
	"fmt"
	"orb-api/config"
	"orb-api/models"
	"orb-api/repositories"
	"orb-api/repositories/task"
	"orb-api/repositories/user"
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
  repo, setupError := config.SetupDB("../.env") 
  
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
