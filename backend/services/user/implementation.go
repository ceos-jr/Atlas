package userservice

import (
	"errors"
	"orb-api/models"
	repository "orb-api/repositories"
	"orb-api/repositories/user"
)

func SetupUserService(repo repository.Repository) *UserService {
	return &UserService{
		UserRepo: repo.User,
	}
}

func (service *UserService) CreateNewUser(credentials ICreateUser) (*models.User, error) {
	// Check if the email is not being used by anyone else
	userArray, readErr := service.UserRepo.ReadBy(user.IReadBy{
		Email: &credentials.Email,
	})

	if readErr != nil {
		return nil, readErr
	}

	if len(userArray) == 1 {
		return nil, errors.New("This email is already being used")
	}

	// Check if the username is not being used by anyone else
	userArray, readErr = service.UserRepo.ReadBy(user.IReadBy{
		Name: &credentials.Name,
	})

	if readErr != nil {
		return nil, readErr
	}

	if len(userArray) == 1 {
		return nil, errors.New("This username is already being used")
	}

	// Check email, username and password length
	if !user.ValidUserName(credentials.Name) {
		return nil, errors.New("Invalid username size")
	}

	if !user.ValidUserEmail(credentials.Email) {
		return nil, errors.New("Invalid email size")
	}

	if !user.ValidUserPassword(credentials.Password) {
		return nil, errors.New("Invalid password size")
	}

	// Hash the password to prevent security vulnerabilities
	hashedPassword, hashErr := HashPassword(credentials.Password)

	if hashErr != nil {
		return nil, hashErr
	}

	newUser, createErr := service.UserRepo.Create(user.ICreate{
		Name:     credentials.Name,
		Email:    credentials.Email,
		Password: hashedPassword,
		Status:   models.UStatusProcessing,
	})

	if createErr != nil {
		return nil, createErr
	}

	return newUser, nil
}

func (service *UserService) UpdateEmail(id uint, Email string) (*string, error) {
	// Checar se o email não já está sendo utilizado
	userArray, readErr := service.UserRepo.ReadBy(user.IReadBy{
		Email: &Email,
	})

	if readErr != nil {
		return nil, readErr
	}

	if len(userArray) == 1 {
		return nil, errors.New("This Email same as current")
	}

	// Checar se o email é diferente do atual
	userArray, readErr = service.UserRepo.ReadBy(user.IReadBy{
		ID:    &id,
		Email: &Email,
	})

	if readErr != nil {
		return nil, readErr
	}

	if len(userArray) == 1 {
		return nil, errors.New("This Email same as current")
	}

	// Atualizando valor
	_, updateErr := service.UserRepo.Update(user.IUpdate{
		ID:    id,
		Email: &Email,
	})

	if updateErr != nil {
		return nil, updateErr
	}

	opa := "deu bom"

	return &opa, nil
}

func (service *UserService) UpdateStatus(id uint, Status uint) (*string, error) {
	// Checar se o status é valido
	if !user.ValidUserStatus(Status) {
		return nil, errors.New("Status Invalido")
	}

	// Checar se o status é diferente do atual
	userStatus, readErr := service.UserRepo.ReadBy(user.IReadBy{
		ID:     &id,
		Status: &Status,
	})

	if readErr != nil {
		return nil, readErr
	}

	if len(userStatus) == 1 {
		return nil, errors.New("This status same as current")
	}

	// Atualizando valor
	_, updateErr := service.UserRepo.Update(user.IUpdate{
		ID:     id,
		Status: &Status,
	})

	if updateErr != nil {
		return nil, updateErr
	}

	opa := "deu bom"

	return &opa, nil

}
