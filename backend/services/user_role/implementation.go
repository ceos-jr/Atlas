package user_role

import (
	"errors"
	"orb-api/models"
	"orb-api/repositories/role"
	"orb-api/repositories/user"
	"orb-api/repositories/user_role"
)

func SetupService(repositoryUser *user.Repository, repositoryRole *role.Repository, repositoryRoleUser *user_role.Repository) *Service {
	return &Service{
		UserRepo: repositoryUser,
		RoleRepo: repositoryRole,
		UserRoleRepo: repositoryRoleUser,

	}
}

func (service *Service) AssigneRole(IdUser uint, IdRole uint) (*models.UserRole, error) {
	if !service.UserRepo.ValidUser(IdUser){
		return nil, errors.New("Invalid user id")
	}

	if !service.RoleRepo.ValidRole(IdRole){
		return nil, errors.New("Invalid role id")
	}

	userArray, readErr := service.UserRepo.ReadBy(user.IReadBy{
		ID: &IdUser,
	})

	if readErr != nil {
		return nil, readErr
	}

	if userArray[0].Status != 2 {
		return nil, errors.New("Invalid user status")
	}

	userroleArray, readErr := service.UserRoleRepo.ReadBy(user_role.IReadBy{
		RoleID: &IdRole,
	})

	if readErr != nil{
		return nil, readErr
	}

	if len(userroleArray) > 0 {
		return nil, errors.New("This role is already assigne")
	}

	newRoleUser, createErr := service.UserRoleRepo.Create(user_role.ICreate{

		UserID:     IdUser,
		RoleID:    IdRole,

	})

	if createErr != nil {
		return nil, createErr
	}

	return newRoleUser, nil
}

func (service *Service) UnassignRole(IdUser uint, IdRole uint) (*models.UserRole, error){
	if !service.UserRepo.ValidUser(IdUser){
		return nil, errors.New("Invalid user id")
	}

	if !service.RoleRepo.ValidRole(IdRole){
		return nil, errors.New("Invalid role id")
	}

	userroleArray, readErr := service.UserRoleRepo.ReadBy(user_role.IReadBy{
		RoleID: &IdRole,
	})

	if userroleArray[0].UserID == nil{
		return nil, errors.New("Role already unassigned")
	}
	
	updateRoleUser, updateErr := service.UserRoleRepo.Update(user_role.IUpdate{
		IdRole: IdRole
		IdUser: nil,
	})

	if updateErr != nil {
		return nil, updateErr
	}

	return updateRoleUser, nil
}