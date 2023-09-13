package role

import (
	"errors"
	"orb-api/models"
	"orb-api/repositories/role"
)

//just to remember the structure of the role model (you can delete this)
//type Role struct {
//ID          uint   `json:"id" gorm:"primaryKey"`
//Name        string `json:"name" gorm:"size:128;not null;"`
//Description string `json:"description" gorm:"not null"`
//}

func (s *ServiceRole) CreateRole(name string, description string ) (*models.Role, error) {
	//logic to create role using roleRepository

	//verify is name or description are not empty; errors.New() creates and returns a new error with the given message.

	if name == "" || description == "" {
		return nil, errors.New("name or description cannot be empty")
	}

	//create a new role using the parameters passed to the function
	newRole := models.Role{
		Name: name,
		Description: description,
}

	//calls the roleRepository to insert the new role in the database
	createdRole, err := s.roleRepository.Create(role.ICreate{
		Name: newRole.Name,
		Description: newRole.Description,
	})

	//if there is an error, return nil and the error
	if err != nil {
		return nil, err
	}

	//if there is no error, return the created role and nil
	return createdRole, nil
}




