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

	//verify is name and description are not empty; errors.New() creates and returns a new error with the given message.

	if name == "" || description == "" {
		return nil, errors.New("name and description cannot be empty")
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


func (s *ServiceRole) GetAllRoles () ([]models.Role, error) {
	//logic to get all roles using roleRepository
}

func (s *ServiceRole) GetRoleByID (id uint) (*models.Role, error) {
	//logic to get role by id using roleRepository
}

func(s *ServiceRole) UpdateRole (id uint, name string, description string) (*models.Role, error) {
	//logic to update role using roleRepository
}

func (s *ServiceRole) DeleteRole (id uint) (*models.Role, error) {
	//logic to delete role using roleRepository
}

func (s *ServiceRole) AssignRoleToUser (roleID uint, userID uint) (*models.User, error) {
	//logic to assign role to user using roleRepository
}

func (s *ServiceRole) RemoveRoleFromUser (roleID uint, userID uint) (*models.User, error) {
	//logic to remove role from user using roleRepository
}

func (s *ServiceRole) GetUsersByRole (roleID uint) ([]models.User, error) {
	//logic to get users by role using roleRepository
}

func (s *ServiceRole) CheckIfUserHasRole (roleID uint, userID uint) (bool, error) {
	//logic to check if user has role using roleRepository
}

func (s *ServiceRole) SearchRoles (name string, description string) ([]models.Role, error) {
	//logic to search roles using roleRepository
}

