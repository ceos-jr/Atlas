package role 

import (
	"orb-api/models"
	"gorm.io/gorm"
)

type (
  Repository struct {
    getDB func () *gorm.DB
  }  
  
  ICreateRole struct {
    Name        string 
    Description string
  } 

  IReadBy struct {
    ID          *uint 
    Name        *string 
    Description *string
  }
  
  IUpdateRole struct {
    RoleID      uint 
    Name        *string 
    Description *string
  }
    
  IDeleteRole struct {
    RoleID      uint
  }

  RoleInterface interface {
    Create(ICreateRole) error 
    ReadAll() ([]models.Role, error)  
    ReadBy(IReadBy) ([]models.Role, error)
    Update(IUpdateRole) error
    Delete(IDeleteRole) error
  }
)
