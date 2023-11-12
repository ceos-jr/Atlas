package user


import (
	"errors"
	"orb-api/models"
	"orb-api/repositories/role"
	"orb-api/repositories/user"
)

func SetupService(repositoryUser *user.Repository, repositoryRole *role.Repository, repositoryRoleUser *user_role.Repository) *Service {
	return &Service{
		UserRepo: repositoryUser,
		RoleRepo: repositoryRole,
		UserRoleRepo: repositoryRoleUser,

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
		Status:   models.UStatusProcessing,we
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

	userroleArray, readErr := service.UserRoleRepo.ReadBy(user.IReadBy{
		RoleID: &IdRole,
	})

	if readErr != nil{
		return nil, readErr
	}

	if len(userroleArray) > 0 {
		return nil, errors.New("This role is already assigne")
	}

	newRoleUser, createErr := service.UserRoleRepo.Create(user.ICreate{

		UserID:     IdUser,
		RoleID:    IdRole,

	})

	if createErr != nil {
		return nil, createErr
	}

	return newRoleUser, nil

}
