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

func (r *UserRoleRepository) Create(createData ICreateUserRole) error {
  var user = models.User { ID: createData.UserID }
  var role = models.Role { ID: createData.RoleID }
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

func (r *UserRoleRepository) ReadAll() ([]models.UserRole, error) {
	var userRoleArray []models.UserRole

	result := r.GetDB().Find(&userRoleArray)

	if result.Error != nil {
		return nil, result.Error
	}

	return userRoleArray, nil
}

func (r *UserRoleRepository) ReadBy(readBy IReadBy) ([]models.UserRole, error) {
	var userRoleArray []models.UserRole
  var userRoleMap map[string]interface{} 
  
  // If the function's caller pass an empty struct as argument 
  // It will return an error
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

func (r *UserRoleRepository) Update(updateData IUpdateUserRole) error {
  var user = models.User{ ID: *updateData.UserID }
  var role = models.Role{ ID: *updateData.RoleID }
  var userRole = models.UserRole{ ID: updateData.UserRoleID }
  var updateMap map[string]interface{}

  // If the function's caller pass an empty struct as argument 
  // It will return an error
  if updateData.UserID == nil && updateData.RoleID == nil {
    return errors.New("No fields to update")
  }
  
  if updateData.UserID != nil {
    updateMap["UserID"] = updateData.UserID
  }

  if updateData.RoleID != nil {
    updateMap["RoleID"] = updateData.RoleID
  }
  
  // Verify userRole existence
	verifyExistence := r.GetDB().First(&userRole)
  
	if verifyExistence.Error != nil {
	  return verifyExistence.Error
	}
  
  // Verify user existence
  verifyUserExistence := r.GetDB().First(&user)

  if verifyUserExistence.Error != nil {
    return verifyUserExistence.Error
  }
  
  // Verify role existence
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

func (r *UserRoleRepository) Delete(deleteData IDeleteUserRole) error {
	var userRole = models.UserRole{ ID: deleteData.UserRoleID }
  
  // verify userRole existence
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
