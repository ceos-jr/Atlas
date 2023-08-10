package repositories

import (
	"orb-api/config"
	"orb-api/models"
)

func CreateRole(repository *config.Repository, name string, description string) (*models.Role, error) {
	newRole := &models.Role{
		Name: name,
		Description: description,
	}

	result := repository.DB.Create(newRole)

	if result.Error != nil {
		return nil, result.Error
	}

	return newRole, nil
}


func GetAllRoles(repository *config.Repository)(allRoles []models.Role){
	var roles []models.Role
	repository.DB.Find(&roles)
	return roles
}

func GetRoleByID(repository *config.Repository, id uint) (*models.Role, error){
	role := &models.Role{ID : id}
	result := repository.DB.Where(&role)

	if result.Error != nil{
		return nil, result.Error
	}

	return role, nil

}


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
