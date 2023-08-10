package userRole

import (
	"errors"
	"gorm.io/gorm"
	"orb-api/models"
	iUserRole "orb-api/repositories/user_role/user_role_interface"
)

type RUserRole struct {
	Repo func() *gorm.DB
}

func (r *RUserRole) Create(newUserRole iUserRole.ICreateUserRole) error {
	var createError error = nil
	var user models.User
	var role models.Role
	var userRole models.UserRole

	verifyExistenceUser := r.Repo().First(&user)

	if verifyExistenceUser.Error != nil {
		createError = verifyExistenceUser.Error
		return createError
	}
	userRole.UserID = newUserRole.UserId

	verifyExistenceRole := r.Repo().First(&role)

	if verifyExistenceRole.Error != nil {
		createError = verifyExistenceRole.Error
		return createError
	}
	userRole.RoleID = newUserRole.RoleId

	result := r.Repo().Create(&userRole)

	if result.Error != nil {
		createError = result.Error
		return createError
	}

	return nil
}

func (r *RUserRole) ReadAll() (*[]models.UserRole, error) {
	var readError error = nil
	var userRoleArray []models.UserRole

	result := r.Repo().Find(&userRoleArray)

	if result.Error != nil {
		readError = result.Error
		return nil, readError
	}

	return &userRoleArray, nil
}

func (r *RUserRole) ReadBy(readBy iUserRole.IReadBy) (*[]models.UserRole, error) {
	var readError error = errors.New("no fields to read")
	var readMap map[string]interface{}
	var userRoleArray []models.UserRole

	if readBy.UserId != nil {
		readError = nil
		readMap["UserId"] = *readBy.UserId
	}

	if readBy.RoleId != nil {
		readError = nil
		readMap["RoleId"] = *readBy.RoleId
	}

	result := r.Repo().Where(readMap).Find(&userRoleArray)
	readError = result.Error

	if readError != nil {
		return nil, readError
	}

	return &userRoleArray, nil
}

func (r *RUserRole) Update(newUserRoleData iUserRole.IUpdateUserRole) error {
	var updateError error = errors.New("no fields to update")
	var updateMap map[string]interface{}
	var userRole models.UserRole = models.UserRole{
		ID: newUserRoleData.UserRoleId,
	}

	verifyExistence := r.Repo().First(&userRole)

	if verifyExistence.Error != nil {
		updateError = verifyExistence.Error
		return updateError
	}

	if newUserRoleData.UserId != nil {
		var user models.User = models.User{
			ID: *newUserRoleData.UserId,
		}

		verifyExistenceUser := r.Repo().First(&user)

		if verifyExistenceUser.Error != nil {
			updateError = verifyExistenceUser.Error
			return updateError
		}

		updateMap["UserID"] = user.ID
	}

	if newUserRoleData.RoleId != nil {
		updateError = nil
		var role models.Role = models.Role{
			ID: *newUserRoleData.RoleId,
		}

		verifyExistenceRole := r.Repo().First(&role)

		if verifyExistenceRole.Error != nil {
			updateError = verifyExistenceRole.Error
			return updateError
		}

		updateMap["RoleID"] = role.ID
	}

	result := r.Repo().Model(&userRole).Updates(updateMap)

	if result.Error != nil {
		updateError = result.Error
		return updateError
	}

	return nil
}

func (r *RUserRole) Delete(user iUserRole.IDeleteUserRole) error {
	var deleteError error
	var userRole models.UserRole = models.UserRole{
		ID: user.UserRoleId,
	}

	verifyExistence := r.Repo().First(&userRole)

	if verifyExistence.Error != nil {
		deleteError = verifyExistence.Error
		return deleteError
	}

	result := r.Repo().Delete(&userRole)

	if result.Error != nil {
		deleteError = result.Error
		return deleteError
	}

	return nil
}
