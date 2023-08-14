package userrole

import (
	"errors"
	"gorm.io/gorm"
	"orb-api/models"
)

func NewUserRoleRepository(connection *gorm.DB) Repository {
	return Repository{
		GetDB: func() *gorm.DB {
			return connection
		},
	}
}

func (r *Repository) Create(createData ICreateUserRole) error {
	var user = models.User{ID: createData.UserID}
	var role = models.Role{ID: createData.RoleID}
	var newUserRole = models.UserRole{
		UserID: createData.UserID,
		RoleID: createData.RoleID,
	}

	verifyUserExistence := r.GetDB().First(&user)

	if verifyUserExistence.Error != nil {
		return verifyUserExistence.Error
	}

	verifyRoleExistence := r.GetDB().First(&role)

	if verifyRoleExistence.Error != nil {
		return verifyRoleExistence.Error
	}

	result := r.GetDB().Create(&newUserRole)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *Repository) ReadAll() ([]models.UserRole, error) {
	var userRoleArray []models.UserRole

	result := r.GetDB().Find(&userRoleArray)

	if result.Error != nil {
		return nil, result.Error
	}

	return userRoleArray, nil
}

func (r *Repository) ReadBy(readBy IReadBy) ([]models.UserRole, error) {
	var userRoleArray []models.UserRole
	var userRoleMap map[string]interface{}

	if readBy.RoleID == nil && readBy.UserID == nil {
		return nil, errors.New("No fields to read")
	}

	if readBy.RoleID != nil {
		userRoleMap["RoleID"] = *readBy.RoleID
	}

	if readBy.UserID != nil {
		userRoleMap["UserID"] = *readBy.UserID
	}

	result := r.GetDB().Where(userRoleMap).Find(&userRoleArray)

	if result.Error != nil {
		return nil, result.Error
	}

	return userRoleArray, nil
}

func (r *Repository) Update(updateData IUpdateUserRole) error {
	var user = models.User{ID: *updateData.UserID}
	var role = models.Role{ID: *updateData.RoleID}
	var userRole = models.UserRole{ID: updateData.UserRoleID}
	var updateMap map[string]interface{}

	if updateData.UserID == nil && updateData.RoleID == nil {
		return errors.New("No fields to update")
	}

	if updateData.UserID != nil {
		updateMap["UserID"] = updateData.UserID
	}

	if updateData.RoleID != nil {
		updateMap["RoleID"] = updateData.RoleID
	}

	verifyExistence := r.GetDB().First(&userRole)

	if verifyExistence.Error != nil {
		return verifyExistence.Error
	}

	verifyUserExistence := r.GetDB().First(&user)

	if verifyUserExistence.Error != nil {
		return verifyUserExistence.Error
	}

	verifyRoleExistence := r.GetDB().First(&role)

	if verifyRoleExistence.Error != nil {
		return verifyRoleExistence.Error
	}

	result := r.GetDB().Model(&userRole).Updates(updateMap)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *Repository) Delete(deleteData IDeleteUserRole) error {
	var userRole = models.UserRole{ID: deleteData.UserRoleID}

	verifyExistence := r.GetDB().First(&userRole)

	if verifyExistence.Error != nil {
		return verifyExistence.Error
	}

	result := r.GetDB().Delete(&userRole)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
