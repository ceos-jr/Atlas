package roleservice

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

func SetupRoleService(repo *role.Repository) *RoleService {
	return &RoleService{
		RoleRepo: repo,
	}
}

func (s *RoleService) CreateRole(name string, description string) (*models.Role, error) {
	//logic to create role using roleRepository

	//verify is name or description are not empty; errors.New() creates and returns a new error with the given message.

	if name == "" || description == "" {
		return nil, errors.New("name or description cannot be empty")
	}

	//create a new role using the parameters passed to the function
	newRole := models.Role{
		Name:        name,
		Description: description,
	}

	//calls the roleRepository to insert the new role in the database
	createdRole, err := s.RoleRepo.Create(role.ICreate{
		Name:        newRole.Name,
		Description: newRole.Description,
	})

	//if there is an error, return nil and the error
	if err != nil {
		return nil, err
	}

	//if there is no error, return the created role and nil
	return createdRole, nil
}

func (s *RoleService) UpdateRoleName(id uint, name string) (*models.Role, error) {
	//check if Role exists
	roleArray, readErr := s.RoleRepo.ReadBy(role.IReadBy{
		ID: &id,
	})

	if readErr != nil {
		return nil, readErr
	}

	if len(roleArray) == 0 {
		return nil, errors.New("This role doesn't exist")
	}

	//check if Name input is null ("")
	if name == "" {
		return nil, errors.New("Name cannot be empty")
	}

	//check if Name already exists
	roleArray, readErr = s.RoleRepo.ReadBy(role.IReadBy{
		Name: &name,
	})
	if len(roleArray) == 1 {
		return nil, errors.New("This name is already being used")
	}

	updateName, updateErr := s.RoleRepo.Update(role.IUpdate{
		RoleID: id,
		Name:   &name,
	})

	if updateErr != nil {
		return nil, updateErr
	}

	return updateName, nil
}

func (s *RoleService) UpdateRoleDescription(id uint, description string) (*models.Role, error) {
	//check if Role exists
	roleArray, readErr := s.RoleRepo.ReadBy(role.IReadBy{
		ID: &id,
	})

	if readErr != nil {
		return nil, readErr
	}

	if len(roleArray) == 0 {
		return nil, errors.New("This role doesn't exist")
	}

	//check if Description is empty / null ("")
	if description == "" {
		return nil, errors.New("Description cannot be empty")
	}

	updateDescription, updateErr := s.RoleRepo.Update(role.IUpdate{
		RoleID:      id,
		Description: &description,
	})

	if updateErr != nil {
		return nil, updateErr
	}

	return updateDescription, nil

}
