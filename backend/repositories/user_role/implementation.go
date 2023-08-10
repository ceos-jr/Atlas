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
	var userRole models.UserRole = models.UserRole{
    UserID: newUserRole.UserID,
    RoleID: newUserRole.RoleID,
  }

	verifyUserExistence := r.GetDB().First(&models.User{
    ID: newUserRole.UserID,
  })

	if verifyUserExistence.Error != nil {
		return verifyUserExistence.Error
	}

	verifyRoleExistence := r.GetDB().First(&models.Role{
    ID: newUserRole.RoleID,
  })

	if verifyRoleExistence.Error != nil {
		return verifyRoleExistence.Error
	}

	result := r.GetDB().Create(&userRole)

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
  var userRole models.UserRole = models.UserRole{
    UserID: readBy.UserID,
    RoleID: readBy.RoleID,
  }
  
  // If the function's caller pass an empty struct as argument 
  // It will return an error
  if readBy.RoleID == 0 && readBy.UserID == 0 {
    return nil, errors.New("No fields to read")
  }	

	result := r.GetDB().Where(&userRole).Find(&userRoleArray)

	if result.Error != nil {
		return nil, result.Error 
	}

	return userRoleArray, nil
}

func (r *UserRoleRepository) Update(updateData IUpdateUserRole) error {
	var userRole models.UserRole = models.UserRole{
		ID: updateData.UserRoleID,
	}
 
  var toUpdate models.UserRole = models.UserRole{
    RoleID: updateData.RoleID,
    UserID: updateData.UserID,
  }

	verifyExistence := r.GetDB().First(&userRole)
  
	if verifyExistence.Error != nil {
	  return verifyExistence.Error
	}
  
  // If the function's caller pass an empty struct as argument 
  // It will return an error
  if updateData.UserID == 0 && updateData.RoleID == 0 {
    return errors.New("No fields to update")
  }
  
  verifyUserExistence := r.GetDB().First(&models.User{
    ID: updateData.UserID,
  })

  if verifyUserExistence.Error != nil {
    return verifyUserExistence.Error
  }

  verifyRoleExistence := r.GetDB().First(&models.Role{
    ID: updateData.RoleID,
  })

  if verifyRoleExistence.Error != nil {
    return verifyRoleExistence.Error
  }

	result := r.GetDB().Model(&userRole).Updates(toUpdate)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *UserRoleRepository) Delete(user IDeleteUserRole) error {
	var userRole models.UserRole = models.UserRole{
		ID: user.UserRoleID,
	}

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
