package user

import (
	"errors"
	"gorm.io/gorm"
	"orb-api/models"
	"time"
)

func NewUserRepository(db *gorm.DB) Repository {
	return Repository{
		GetDB: func() *gorm.DB {
			return db
		},
	}
}

func ValidUserName(name string) bool {
	if len(name) < nameMinLen || len(name) > nameMaxLen {
		return false
	}
	return true
}

func ValidUserEmail(email string) bool {
	if len(email) < emailMinLen || len(email) > emailMaxLen {
		return false
	}
	return true
}

func ValidUserPassword(password string) bool {
	if len(password) < passwordMinLen {
		return false
	}
	return true
}

func ValidUserStatus(status uint) bool {
	_, valid := models.UserStatus[status]

	return valid
}

func (r *Repository) ValidUser(id uint) bool {
	user := models.User{ID: id}

	verifyUser := r.GetDB().First(&user).Error

	if verifyUser != nil {
		return false
	}

	return true
}

func (r *Repository) Create(createData ICreate) (*models.User, error) {
	var user = models.User{
		Name:      createData.Name,
		Email:     createData.Email,
		Password:  createData.Password,
		Status:    createData.Status,
		UpdatedAt: time.Now(),
	}

	if !ValidUserEmail(createData.Email) {
		return nil, errors.New("invalid email value")
	}

	if !ValidUserName(createData.Name) {
		return nil, errors.New("invalid name value")
	}

	if !ValidUserPassword(createData.Password) {
		return nil, errors.New("invalid password value")
	}

	if !ValidUserStatus(createData.Status) {
		return nil, errors.New("invalid status")
	}

	result := r.GetDB().Create(&user)

	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

func (r *Repository) ReadAll(all IReadAll) ([]models.User, error) {
	var result *gorm.DB
	var userArray []models.User

	if all.Limit != nil {
		result = r.GetDB().Find(&userArray).Limit(*all.Limit)
	} else {
		result = r.GetDB().Find(&userArray)
	}

	if result.Error != nil {
		return nil, result.Error
	}

	return userArray, nil
}

func (r *Repository) ReadBy(readBy IReadBy) ([]models.User, error) {
	var fieldMap = make(map[string]interface{})
	var userArray []models.User
	var result *gorm.DB

	if readBy.ID == nil &&
		readBy.Status == nil &&
		readBy.Email == nil &&
		readBy.Name == nil {
		return nil, errors.New("No fields to read")
	}

	if readBy.ID != nil {
		fieldMap["id"] = *readBy.ID
	}

	if readBy.Name != nil {
		if !ValidUserName(*readBy.Name) {
			return nil, errors.New("Invalid name")
		}

		fieldMap["name"] = *readBy.Name
	}

	if readBy.Email != nil {
		if !ValidUserName(*readBy.Email) {
			return nil, errors.New("Invalid email")
		}

		fieldMap["email"] = *readBy.Email
	}

	if readBy.Status != nil {
		if !ValidUserStatus(*readBy.Status) {
			return nil, errors.New("Invalid status")
		}

		fieldMap["status"] = *readBy.Status
	}

	if readBy.Limit != nil {
		result = r.GetDB().Where(fieldMap).Find(&userArray).Limit(*readBy.Limit)
	} else {
		result = r.GetDB().Where(fieldMap).Find(&userArray)
	}

	if result.Error != nil {
		return nil, result.Error
	}

	return userArray, nil
}

func (r *Repository) Update(updateData IUpdate) (*models.User, error) {
	var fieldMap = make(map[string]interface{})
	var user = models.User{ID: updateData.ID}

	if updateData.Name == nil &&
		updateData.Email == nil &&
		updateData.Status == nil &&
		updateData.Password == nil {
		return nil, errors.New("No fields to update")
	}

	if updateData.Name != nil {
		if !ValidUserName(*updateData.Name) {
			return nil, errors.New("Invalid name")
		}

		fieldMap["name"] = *updateData.Name
	}

	if updateData.Email != nil {
		if !ValidUserEmail(*updateData.Email) {
			return nil, errors.New("Invalid email")
		}

		fieldMap["email"] = *updateData.Email
	}

	if updateData.Status != nil {
		if !ValidUserStatus(*updateData.Status) {
			return nil, errors.New("Invalid status")
		}

		fieldMap["status"] = *updateData.Status
	}

	if updateData.Password != nil {
		if !ValidUserPassword(*updateData.Password) {
			return nil, errors.New("Invalid password")
		}

		fieldMap["password"] = *updateData.Password
	}

	fieldMap["updated_at"] = time.Now()

	result := r.GetDB().Model(&user).Updates(fieldMap)

	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

func (r *Repository) Delete(deleteData IDelete) (*models.User, error) {
	var user = models.User{ID: deleteData.ID}

	verifyExistence := r.GetDB().First(&user)

	if verifyExistence.Error != nil {
		return nil, verifyExistence.Error
	}

	result := r.GetDB().Delete(&user)

	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}
