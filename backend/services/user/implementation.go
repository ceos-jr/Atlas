package user


import (
	"errors"
	"orb-api/models"
	"orb-api/repositories/user"
	"orb-api/repositories/userproject"
	"orb-api/repositories/project"
)

func SetupService(repository1 *user.Repository, repository2 *userproject.Repository, repository3 *project.Repository) *Service {
	return &Service{
		UserRepo: 			repository1,
		UserProjectRepo:	repository2,
		ProjectRepo:		repository3,
	}
}

func (service *Service) CreateUser(
	name string,
	email string,
	//password string,
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

	/*if !user.ValidUserPassword(password) {

		return nil, errors.New("Invalid password size")
	}

	// Hash the password to prevent security vulnerabilities

	hashedPassword, hashErr := HashPassword(password)


	if hashErr != nil {
		return nil, hashErr
	}*/

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

func (service *Service) DeleteUser(id uint) (*models.User, error) {
	if !service.UserRepo.ValidUser(id) {
		return nil, errors.New("Invalid user id")
	}

	userArray, readErr := service.UserRepo.ReadBy(user.IReadBy{
		ID: &id,
	})

	if readErr != nil {
		return nil, readErr
	}

	if userArray[0].Status == uint(1) {
		return nil, errors.New("User already disabled")
	}
	
	status := uint(1)

	userUpdate, updateErr := service.UserRepo.Update(user.IUpdate{
		ID:     id,
		Status: &status,
	})

	if updateErr != nil {
		return nil, updateErr
	}

	return userUpdate, nil

}
func (service *Service) SortProjects(UserID uint) ([]models.Project, error){
	UserProjects, ReadErr := service.UserProjectRepo.ReadBy(userproject.IReadBy{
		UserID: &UserID,
	})

	if ReadErr != nil {
		return nil, ReadErr
	}

	Projects := []models.Project{}

	for i := 0; i < len(UserProjects); i++{
		App, Err := service.ProjectRepo.ReadBy(project.IReadBy{
			ID: &(UserProjects[i].ProjectID),
		})

		if Err != nil {
			return nil, Err
		}

		Projects = append(Projects, App...)
	}

	return Projects, nil

}

func (service *Service) ReadUser(read user.IReadBy) ([]models.User, error) {
	if  read.ID == nil &&
		read.Name == nil &&
		read.Email == nil &&
		read.Status == nil &&
		read.Limit == nil {
		usersArray, readErr := service.UserRepo.ReadAll(user.IReadAll{})

		if readErr != nil{
			return nil, readErr
		}
		return usersArray, nil
	} 

	usersArray, readErr := service.UserRepo.ReadBy(read)
	
	if readErr != nil{
		return nil, readErr
	}
	return usersArray, nil
}
