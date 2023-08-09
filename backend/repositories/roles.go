package repositories

import (
	"orb-api/config"
	"orb-api/models"
)

func GetRoleByName(repository *config.Repository, name string) (*models.Role, error) {

	role := &models.Role{Name: name}

	result := repository.DB.First(role)

	if result.Error != nil {
		return nil, result.Error
	}

	return role, nil
}

func DeleteRole(repository *config.Repository, roleID uint) error {

	if err := repository.DB.Delete(&models.Role{}, roleID).Error; err != nil {
		return err
	}
	return nil

}

//This function only updates non-zero fields as default

func UpdateRole(repository *config.Repository, name, description string, roleID uint) error {
	role := &models.Role{
		ID: roleID,
	}

	updateError := repository.DB.Model(&role).Updates(&models.Role{Name: name, Description: description}).Error

	if updateError != nil {
		return updateError
	}
	return nil
}
