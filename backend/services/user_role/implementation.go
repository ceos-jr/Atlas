package user_role

import (
	"errors"
	"orb-api/models"
	"orb-api/repositories/role"
	"orb-api/repositories/user"
	"orb-api/repositories/user_role"
)

func SetupService(repositoryUser *user.Repository, repositoryRole *role.Repository, repositoryRoleUser *userrole.Repository) *Service {
	return &Service{
		UserRepo: repositoryUser,
		RoleRepo: repositoryRole,
		UserRoleRepo: repositoryRoleUser,
	}
}

func (service *Service) AssigneRole(IdUser uint, IdRole uint) (*models.UserRole, error) {
	if !service.UserRepo.ValidUser(IdUser){
		return nil, errors.New("invalid user id")
	}

	if !service.RoleRepo.ValidRole(IdRole){
		return nil, errors.New("invalid role id")
	}

	userArray, readErr := service.UserRepo.ReadBy(user.IReadBy{
		ID: &IdUser,
	})

	if readErr != nil {
		return nil, readErr
	}

	if userArray[0].Status != 2 {
		return nil, errors.New("invalid user status")
	}

	userroleArray, readErr := service.UserRoleRepo.ReadBy(userrole.IReadBy{
		RoleID: &IdRole,
	})

	if readErr != nil{
		return nil, readErr
	}

	if len(userroleArray) != 0 {//Ponto a ser analisado sobre não ter duas pessoas com o mesmo cargo
		return nil, errors.New("this role is already assigne")
	}

	newRoleUser, createErr := service.UserRoleRepo.Create(userrole.ICreate{

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
		return nil, errors.New("invalid user id")
	}

	if !service.RoleRepo.ValidRole(IdRole){
		return nil, errors.New("invalid role id")
	}

	userroleArray, readErr := service.UserRoleRepo.ReadBy(userrole.IReadBy{
		UserID: &IdUser,
		RoleID: &IdRole,
	})

	if readErr != nil{
		return nil, readErr
	}

	IDuserrole := userroleArray[0].ID

	deletedUserRole, deleteErr := service.UserRoleRepo.Delete(userrole.IDelete{
		UserRoleID: IDuserrole,
	})

	if deleteErr != nil {
		return nil, deleteErr
	}
	
	return deletedUserRole, nil
}