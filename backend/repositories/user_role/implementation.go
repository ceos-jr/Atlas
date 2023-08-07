package userRoleRepository

import (
	"gorm.io/gorm"
	"orb-api/models"
)

type UserRepository struct {
	db *gorm.DB
}

func (i UserRepository) ReadAll() IReadResultUserRole {
	var userRoleArray []models.UserRole
	result := i.db.Find(&userRoleArray)
	return IReadResultUserRole{
		data:   &userRoleArray,
		status: result.Error,
	}
}

func (i UserRepository) ReadByUser(readUser IReadByUser) IReadResultUserRole {
	var userRoleArray []models.UserRole
	result := i.db.Where(
		map[string]interface{}{
			"id": readUser.userId,
		}).Find(&userRoleArray)

	return IReadResultUserRole{
		data:   &userRoleArray,
		status: result.Error,
	}
}
