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

func (r *Repository) Create(createData ICreate) error {
	var user = models.User{
		Name:      createData.Name,
		Email:     createData.Email,
		Password:  createData.Password,
		Status:    createData.Status,
		UpdatedAt: time.Now(),
	}

	if !ValidUserEmail(createData.Email) {
		return errors.New("invalid email value")
	}

	if !ValidUserName(createData.Name) {
		return errors.New("invalid name value")
	}

	if !ValidUserPassword(createData.Password) {
		return errors.New("invalid password value")
	}

	if !ValidUserStatus(createData.Status) {
		return errors.New("invalid status")
	}

	result := r.GetDB().Create(&user)

	if result.Error != nil {
		return result.Error
	}

	return nil
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
	var fieldMap map[string]interface{}
	var userArray []models.User
  var result *gorm.DB

	if readBy.ID == nil &&
		readBy.Status == nil &&
		readBy.Email == nil &&
		readBy.Name == nil {
		return nil, errors.New("No fields to read")
	}
  
  if readBy.ID != nil {
    fieldMap["ID"] = *readBy.ID
  }

  if readBy.Name != nil {
    if !ValidUserName(*readBy.Name) {
      return nil, errors.New("Invalid name")
    }

    fieldMap["Name"] = *readBy.Name
  }

  if readBy.Email != nil {
    if !ValidUserName(*readBy.Email) {
      return nil, errors.New("Invalid email")
    }

    fieldMap["Email"] = *readBy.Email
  }

  if readBy.Status != nil {
    if !ValidUserStatus(*readBy.Status) {
      return nil, errors.New("Invalid status")
    }

    fieldMap["Status"] = *readBy.Status
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

func (r *Repository) Update(updateData IUpdate) error {
	var fieldMap map[string]interface{}
	var user = models.User{ID: updateData.ID}
  
  if updateData.Name == nil && 
    updateData.Email == nil && 
    updateData.Status == nil {
    return errors.New("No fields to update")
  } 

  if updateData.Name != nil {
    if !ValidUserName(*updateData.Name) {
      return errors.New("Invalid name")
    }

    fieldMap["Name"] = *updateData.Name
  }

  if updateData.Email != nil {
    if !ValidUserName(*updateData.Email) {
      return errors.New("Invalid email")
    }

    fieldMap["Email"] = *updateData.Email
  }

  if updateData.Status != nil {
    if !ValidUserStatus(*updateData.Status) {
      return errors.New("Invalid status")
    }

    fieldMap["Status"] = *updateData.Status
  }
  
  fieldMap["UpdatedAt"] = time.Now()

  result := r.GetDB().Model(&user).Updates(fieldMap)

  if result.Error != nil {
    return result.Error
  }

  return nil
}

func (r *Repository) Delete(deleteData IDelete) error {
	var user = models.User{ID: deleteData.ID}

	verifyExistence := r.GetDB().First(&user)

	if verifyExistence.Error != nil {
		return verifyExistence.Error
	}

  result := r.GetDB().Delete(&user)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
