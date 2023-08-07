package UserRolesRepo

import (
	"orb-api/models"
	"orb-api/repositories"
)

type RUserRole struct {
	repo *repositories.Repository
}

func Setup(repo *repositories.Repository) RUserRole {
	return RUserRole{
		repo: repo,
	}
}

func (r RUserRole) ReadAll() (*[]models.UserRole, error) {
	var readError error
	var userRoleArray []models.UserRole

	readError = nil
	result := r.repo.DB.Find(&userRoleArray)

	if result.Error != nil {
		readError = result.Error
	}

	return &userRoleArray, readError
}

func (r RUserRole) ReadByUser(readUser IReadByUser) (*[]models.UserRole, error) {
	var readError error
	var userRoleArray []models.UserRole

	readError = nil
	result := r.repo.DB.Where(
		map[string]interface{}{
			"id": readUser.UserId,
		}).Find(&userRoleArray)

	if result.Error != nil {
		readError = result.Error
	}

	return &userRoleArray, readError
}
