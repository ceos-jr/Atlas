package userRole

import (
	"errors"
	"orb-api/models"
	"orb-api/repositories"
)

func Setup(repo *repositories.Repository) RUserRole {
	return RUserRole{
		repo: repo,
	}
}

func (r *RUserRole) Create(newUserRole ICreateUserRole) error {
	var createError error = nil
	var user models.User
	var role models.Role
	var userRole models.UserRole

	verifyExistenceUser := r.repo.DB.First(&user)

	if verifyExistenceUser.Error != nil {
		createError = verifyExistenceUser.Error
		return createError
	}
	userRole.UserID = newUserRole.UserId

	verifyExistenceRole := r.repo.DB.First(&role)

	if verifyExistenceRole.Error != nil {
		createError = verifyExistenceRole.Error
		return createError
	}
	userRole.RoleID = newUserRole.RoleId

	result := r.repo.DB.Create(&userRole)

	if result.Error != nil {
		createError = result.Error
		return createError
	}

	return nil
}

func (r *RUserRole) ReadAll() (*[]models.UserRole, error) {
	var readError error = nil
	var userRoleArray []models.UserRole

	result := r.repo.DB.Find(&userRoleArray)

	if result.Error != nil {
		readError = result.Error
		return nil, readError
	}

	return &userRoleArray, nil
}

func (r *RUserRole) ReadBy(readBy IReadBy) (*[]models.UserRole, error) {
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

	result := r.repo.DB.Where(readMap).Find(&userRoleArray)
	readError = result.Error

	if readError != nil {
		return nil, readError
	}

	return &userRoleArray, nil
}

func (r *RUserRole) Update(newUserRoleData IUpdateUserRole) error {
	var updateError error = errors.New("no fields to update")
	var updateMap map[string]interface{}
	var userRole models.UserRole = models.UserRole{
		ID: newUserRoleData.UserRoleId,
	}

	verifyExistence := r.repo.DB.First(&userRole)

	if verifyExistence.Error != nil {
		updateError = verifyExistence.Error
		return updateError
	}

	if newUserRoleData.newUserId != nil {
		var user models.User = models.User{
			ID: *newUserRoleData.newUserId,
		}

		verifyExistenceUser := r.repo.DB.First(&user)

		if verifyExistenceUser.Error != nil {
			updateError = verifyExistenceUser.Error
			return updateError
		}

		updateMap["UserID"] = user.ID
	}

	if newUserRoleData.newUserId != nil {
		updateError = nil
		var role models.Role = models.Role{
			ID: *newUserRoleData.newRoleId,
		}

		verifyExistenceRole := r.repo.DB.First(&role)

		if verifyExistenceRole.Error != nil {
			updateError = verifyExistenceRole.Error
			return updateError
		}

		updateMap["RoleID"] = role.ID
	}

	result := r.repo.DB.Model(&userRole).Updates(updateMap)

	if result.Error != nil {
		updateError = result.Error
		return updateError
	}

	return nil
}

func (r *RUserRole) Delete(user IDeleteUserRole) error {
	var deleteError error
	var userRole models.UserRole = models.UserRole{
		ID: user.UserRoleId,
	}

	verifyExistence := r.repo.DB.First(&userRole)

	if verifyExistence.Error != nil {
		deleteError = verifyExistence.Error
		return deleteError
	}

	result := r.repo.DB.Delete(&userRole)

	if result.Error != nil {
		deleteError = result.Error
		return deleteError
	}

	return nil
}
