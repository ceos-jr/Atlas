package userrole

import (
	"errors"
	"orb-api/models"

	"gorm.io/gorm"
)

func NewUserRoleRepository(connection *gorm.DB) Repository {
	return Repository{
		GetDB: func() *gorm.DB {
			return connection
		},
	}
}

func (r *Repository) Create(createData ICreate) (*models.UserRole, error) {
	var user = models.User{ID: createData.UserID}
	var role = models.Role{ID: createData.RoleID}
	var userRole = models.UserRole{
		UserID: createData.UserID,
		RoleID: createData.RoleID,
	}

	verifyUserExistence := r.GetDB().First(&user)

	if verifyUserExistence.Error != nil {
		return nil, verifyUserExistence.Error
	}

	verifyRoleExistence := r.GetDB().First(&role)

	if verifyRoleExistence.Error != nil {
		return nil, verifyRoleExistence.Error
	}

	result := r.GetDB().Create(&userRole)

	if result.Error != nil {
		return nil, result.Error
	}

	return &userRole, nil
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
	var userRoleMap = make(map[string]interface{})

	if readBy.RoleID == nil && readBy.UserID == nil {
		return nil, errors.New("No fields to read")
	}

	if readBy.RoleID != nil {
		userRoleMap["role_id"] = *readBy.RoleID
	}

	if readBy.UserID != nil {
		userRoleMap["user_id"] = *readBy.UserID
	}

	result := r.GetDB().Where(userRoleMap).Find(&userRoleArray)

	if result.Error != nil {
		return nil, result.Error
	}

	return userRoleArray, nil
}

func (r *Repository) Update(updateData IUpdate) (*models.UserRole, error) {
	userRole := models.UserRole{ID: updateData.UserRoleID}
	var updateMap = make(map[string]interface{})

	if updateData.UserID == nil && updateData.RoleID == nil {
		return nil, errors.New("No fields to update")
	}

	if updateData.UserID != nil {
		updateMap["user_id"] = *updateData.UserID
		user := models.User{ID: *updateData.UserID}
		verifyUserExistence := r.GetDB().First(&user)

		if verifyUserExistence.Error != nil {
			return nil, verifyUserExistence.Error
		}
	}

	if updateData.RoleID != nil {
		updateMap["role_id"] = *updateData.RoleID
		role := models.Role{ID: *updateData.RoleID}
		verifyRoleExistence := r.GetDB().First(&role)

		if verifyRoleExistence.Error != nil {
			return nil, verifyRoleExistence.Error
		}
	}

	result := r.GetDB().Model(&userRole).Updates(updateMap)

	if result.Error != nil {
		return nil, result.Error
	}

	return &userRole, nil
}

func (r *Repository) Delete(deleteData IDelete) (*models.UserRole, error) {
	var userRole = models.UserRole{ID: deleteData.UserRoleID}

	verifyExistence := r.GetDB().First(&userRole)

	if verifyExistence.Error != nil {
		return nil, verifyExistence.Error
	}

	result := r.GetDB().Delete(&userRole)

	if result.Error != nil {
		return nil, result.Error
	}

	return &userRole, nil
}
