package userrole

import (
	"errors"
	"gorm.io/gorm"
	"orb-api/models"
)

func NewUserRoleRepository(connection *gorm.DB) UserRoleRepository {
  return UserRoleRepository{
    GetDB: func () *gorm.DB{
      return connection
    },  
  }
}

func (r *UserRoleRepository) Create(newUserRole ICreateUserRole) error {
	var createError error = nil
	var user models.User
	var role models.Role
	var userRole models.UserRole

	verifyExistenceUser := r.GetDB().First(&user)

	if verifyExistenceUser.Error != nil {
		createError = verifyExistenceUser.Error
		return createError
	}
	userRole.UserID = newUserRole.UserID

	verifyExistenceRole := r.GetDB().First(&role)

	if verifyExistenceRole.Error != nil {
		createError = verifyExistenceRole.Error
		return createError
	}
	userRole.RoleID = newUserRole.RoleID

	result := r.GetDB().Create(&userRole)

	if result.Error != nil {
		createError = result.Error
		return createError
	}

	return nil
}

func (r *UserRoleRepository) ReadAll() (*[]models.UserRole, error) {
	var readError error = nil
	var userRoleArray []models.UserRole

	result := r.GetDB().Find(&userRoleArray)

	if result.Error != nil {
		readError = result.Error
		return nil, readError
	}

	return &userRoleArray, nil
}

func (r *UserRoleRepository) ReadBy(readBy IReadBy) (*[]models.UserRole, error) {
	var readError error = errors.New("no fields to read")
	var readMap map[string]interface{}
	var userRoleArray []models.UserRole

	if readBy.UserID != nil {
		readError = nil
		readMap["UserId"] = *readBy.UserID
	}

	if readBy.RoleID != nil {
		readError = nil
		readMap["RoleId"] = *readBy.RoleID
	}

	result := r.GetDB().Where(readMap).Find(&userRoleArray)
	readError = result.Error

	if readError != nil {
		return nil, readError
	}

	return &userRoleArray, nil
}

func (r *UserRoleRepository) Update(newUserRoleData IUpdateUserRole) error {
	var updateError error = errors.New("no fields to update")
	var updateMap map[string]interface{}
	var userRole models.UserRole = models.UserRole{
		ID: newUserRoleData.UserRoleID,
	}

	verifyExistence := r.GetDB().First(&userRole)

	if verifyExistence.Error != nil {
		updateError = verifyExistence.Error
		return updateError
	}

	if newUserRoleData.UserID != nil {
		var user models.User = models.User{
			ID: *newUserRoleData.UserID,
		}

		verifyExistenceUser := r.GetDB().First(&user)

		if verifyExistenceUser.Error != nil {
			updateError = verifyExistenceUser.Error
			return updateError
		}

		updateMap["UserID"] = user.ID
	}

	if newUserRoleData.RoleID != nil {
		updateError = nil
		var role models.Role = models.Role{
			ID: *newUserRoleData.RoleID,
		}

		verifyExistenceRole := r.GetDB().First(&role)

		if verifyExistenceRole.Error != nil {
			updateError = verifyExistenceRole.Error
			return updateError
		}

		updateMap["RoleID"] = role.ID
	}

	result := r.GetDB().Model(&userRole).Updates(updateMap)

	if result.Error != nil {
		updateError = result.Error
		return updateError
	}

	return nil
}

func (r *UserRoleRepository) Delete(user IDeleteUserRole) error {
	var deleteError error
	var userRole models.UserRole = models.UserRole{
		ID: user.UserRoleId,
	}

	verifyExistence := r.GetDB().First(&userRole)

	if verifyExistence.Error != nil {
		deleteError = verifyExistence.Error
		return deleteError
	}

	result := r.GetDB().Delete(&userRole)

	if result.Error != nil {
		deleteError = result.Error
		return deleteError
	}

	return nil
}
