package user

import (
	"errors"
	"orb-api/models"
	"orb-api/repositories/user"
)

func SetupService(repository *user.Repository) *Service {
	return &Service{
		UserRepo: repository,
	}
}

func (service *Service) NewUser(
	name string,
	email string,
	password string,
) (*models.User, error) {
	// Check if the email is not being used by anyone else
	userArray, readErr := service.UserRepo.ReadBy(user.IReadBy{
		Email: &email,
	})

	if readErr != nil {
		return nil, readErr
	}

	if len(userArray) == 1 {
		return nil, errors.New("This email is already being used")
	}

	// Check if the username is not being used by anyone else
	userArray, readErr = service.UserRepo.ReadBy(user.IReadBy{
		Name: &name,
	})

	if readErr != nil {
		return nil, readErr
	}

	if len(userArray) == 1 {
		return nil, errors.New("This username is already being used")
	}

	// Check email, username and password length
	if !user.ValidUserName(name) {
		return nil, errors.New("Invalid username size")
	}

	if !user.ValidUserEmail(email) {
		return nil, errors.New("Invalid email size")
	}

	if !user.ValidUserPassword(password) {
		return nil, errors.New("Invalid password size")
	}

	// Hash the password to prevent security vulnerabilities
	hashedPassword, hashErr := HashPassword(password)

	if hashErr != nil {
		return nil, hashErr
	}

	newUser, createErr := service.UserRepo.Create(user.ICreate{
		Name:     name,
		Email:    email,
		Password: hashedPassword,
		Status:   models.UStatusProcessing,
	})

	if createErr != nil {
		return nil, createErr
	}

	return newUser, nil
}

func (service *Service) UpdateName(id uint, name string) (*models.User, error) {
	// Check if name has a valid length
	if !user.ValidUserName(name) {
		return nil, errors.New("Invalid username size")
	}

	// Check if the id belongs a valid user
	if !service.UserRepo.ValidUser(id) {
		return nil, errors.New("Invalid user id")
	}

	// Check if the name is not being used by anyone else and different by current
	userArray, readErr := service.UserRepo.ReadBy(user.IReadBy{
		Name: &name,
	})

	if readErr != nil {
		return nil, readErr
	}

	if len(userArray) == 1 {
		return nil, errors.New("This name is already being used")
	}

	// Update username
	updatedUser, updateErr := service.UserRepo.Update(user.IUpdate{
		ID:   id,
		Name: &name,
	})

	if updateErr != nil {
		return nil, updateErr
	}

	return updatedUser, nil
}

func (service *Service) UpdatePassword(id uint, password string) (*models.User, error) {
	// Check if it is password has a valid length
	if !user.ValidUserPassword(password) {
		return nil, errors.New("Invalid password size")
	}

	// Check if the id belongs a valid user
	if !service.UserRepo.ValidUser(id) {
		return nil, errors.New("Invalid user id")
	}

	// hash password
	hashedPassword, hashError := HashPassword(password)

	if hashError != nil {
		return nil, hashError
	}

	updatedUser, updateError := service.UserRepo.Update(user.IUpdate{
		ID:       id,
		Password: &hashedPassword,
	})

	if updateError != nil {
		return nil, updateError
	}

	return updatedUser, nil
}

func (service *Service) UpdateEmail(id uint, email string) (*models.User, error) {
	if !user.ValidUserEmail(email) {
		return nil, errors.New("Invalid email size")
	}

	// Check if the id belongs a valid user
	if !service.UserRepo.ValidUser(id) {
		return nil, errors.New("Invalid user id")
	}

	// Check if the email is not being used by anyone else and different by current
	userArray, readErr := service.UserRepo.ReadBy(user.IReadBy{
		Email: &email,
	})

	if readErr != nil {
		return nil, readErr
	}

	if len(userArray) == 1 {
		return nil, errors.New("This email is already being used")
	}

	// Update e-mail
	userUpdate, updateErr := service.UserRepo.Update(user.IUpdate{
		ID:    id,
		Email: &email,
	})

	if updateErr != nil {
		return nil, updateErr
	}

	return userUpdate, nil
}

func (service *Service) UpdateStatus(id uint, status uint) (*models.User, error) {
	// Check if the status is valid
	if !user.ValidUserStatus(status) {
		return nil, errors.New("Invalid status")
	}

	// Check if the id belongs a valid user
	if !service.UserRepo.ValidUser(id) {
		return nil, errors.New("Invalid user id")
	}

	// Update status
	userUpdate, updateErr := service.UserRepo.Update(user.IUpdate{
		ID:     id,
		Status: &status,
	})

	if updateErr != nil {
		return nil, updateErr
	}

	return userUpdate, nil
}

func (service *Service) HelperUpdateName(id uint, name string) error {
	// Check if name has a valid length
	if !user.ValidUserName(name) {
		return errors.New("Invalid username size")
	}

	// Check if the id belongs a valid user
	if !service.UserRepo.ValidUser(id) {
		return errors.New("Invalid user id")
	}

	// Check if the name is not being used by anyone else and different by current
	userArray, readErr := service.UserRepo.ReadBy(user.IReadBy{
		Name: &name,
	})

	if readErr != nil {
		return readErr
	}

	if len(userArray) == 1 {
		return errors.New("This name is already being used")
	}

	return nil
}

func (service *Service) HelperUpdatePassword(id uint, password string) error {
	// Check if it is password has a valid length
	if !user.ValidUserPassword(password) {
		return errors.New("Invalid password size")
	}

	// Check if the id belongs a valid user
	if !service.UserRepo.ValidUser(id) {
		return errors.New("Invalid user id")
	}

	// hash password
	_, hashError := HashPassword(password)

	if hashError != nil {
		return hashError
	}

	return nil
}

func (service *Service) HelperUpdateEmail(id uint, email string) error {
	if !user.ValidUserEmail(email) {
		return errors.New("Invalid email size")
	}

	// Check if the id belongs a valid user
	if !service.UserRepo.ValidUser(id) {
		return errors.New("Invalid user id")
	}

	// Check if the email is not being used by anyone else and different by current
	userArray, readErr := service.UserRepo.ReadBy(user.IReadBy{
		Email: &email,
	})

	if readErr != nil {
		return readErr
	}

	if len(userArray) == 1 {
		return errors.New("This email is already being used")
	}

	return nil
}

func (service *Service) HelperUpdateStatus(id uint, status uint) error {
	// Check if the status is valid
	if !user.ValidUserStatus(status) {
		return errors.New("Invalid status")
	}

	// Check if the id belongs a valid user
	if !service.UserRepo.ValidUser(id) {
		return errors.New("Invalid user id")
	}

	// Update status

	return nil
}

func (service *Service) UpdateUser(updateData user.IUpdate) (*models.User, error) {
	var user = user.IUpdate{ID: updateData.ID}
	updateID := updateData.ID

	updateNameErr := service.HelperUpdateName(updateID, *updateData.Name)
	if updateNameErr == nil {
		user.Name = updateData.Name
	}

	updateEmailErr := service.HelperUpdateEmail(updateID, *updateData.Email)
	if updateEmailErr == nil {
		user.Email = updateData.Email
	}

	updatePasswordErr := service.HelperUpdatePassword(updateID, *updateData.Password)
	if updatePasswordErr == nil {
		user.Password = updateData.Password
	}

	updateStatusErr := service.HelperUpdateStatus(updateID, *updateData.Status)
	if updateStatusErr == nil {
		user.Status = updateData.Status
	}

	updatedUser, updateErr := service.UserRepo.Update(user)

	if updateErr != nil {
		return nil, updateErr
	}

	return updatedUser, nil
}
