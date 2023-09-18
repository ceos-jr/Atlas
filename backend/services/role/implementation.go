package role

import (
	"errors"
	"orb-api/models"
	"orb-api/repositories/role"
)

func SetupService(repo *role.Repository) *Service {
	return &Service{
		RoleRepo: repo,
	}
}

func (s *Service) CreateRole(name string, description string) (*models.Role, error) {
	if name == "" || description == "" {
		return nil, errors.New("name or description cannot be empty")
	}

	createdRole, err := s.RoleRepo.Create(role.ICreate{
		Name:        name,
		Description: description,
	})

	if err != nil {
		return nil, err
	}

	return createdRole, nil
}

func (s *Service) UpdateName(id uint, name string) (*models.Role, error) {
	if name == "" {
		return nil, errors.New("Name cannot be empty")
	}

	if !s.RoleRepo.ValidRole(id) {
		return nil, errors.New("This role doesn't exist")
	}

	roleArray, readErr := s.RoleRepo.ReadBy(role.IReadBy{
		Name: &name,
	})

	if readErr != nil {
		return nil, readErr
	}

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

func (s *Service) UpdateDescription(id uint, description string) (*models.Role, error) {
	if description == "" {
		return nil, errors.New("Description cannot be empty")
	}

	if !s.RoleRepo.ValidRole(id) {
		return nil, errors.New("This role doesn't exist")
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
