package role

import (
	"errors"
	"orb-api/models"

	"gorm.io/gorm"
)

func NewRoleRepository(connection *gorm.DB) Repository {
	return Repository{
		getDB: func() *gorm.DB {
			return connection
		},
	}
}

func (r *Repository) Create(createData ICreate) (*models.Role, error) {
	var newRole = models.Role{
		Name:        createData.Name,
		Description: createData.Description,
	}

	result := r.getDB().Create(&newRole)

	if result.Error != nil {
		return nil, result.Error
	}

	return &newRole, nil
}

func (r *Repository) ReadAll(all IReadAll) ([]models.Role, error) {
	var rolesArray []models.Role
	var result *gorm.DB

	if all.Limit != nil {
		result = r.getDB().Find(&rolesArray).Limit(*all.Limit)
	} else {
		result = r.getDB().Find(&rolesArray)
	}

	if result.Error != nil {
		return nil, result.Error
	}

	return rolesArray, nil
}

func (r *Repository) ReadBy(readBy IReadBy) ([]models.Role, error) {
	var rolesArray []models.Role
	var roleMap = make(map[string]interface{})

	if readBy.ID == nil && readBy.Name == nil && readBy.Description == nil {
		return nil, errors.New("No fields to read")
	}

	if readBy.ID != nil {
		roleMap["ID"] = readBy.ID
	}

	if readBy.Name != nil {
		roleMap["Name"] = readBy.Name
	}

	if readBy.Description != nil {
		roleMap["Description"] = readBy.Description
	}

	result := r.getDB().Where(roleMap).Find(&rolesArray)

	if result.Error != nil {
		return nil, result.Error
	}

	return rolesArray, nil
}

func (r *Repository) Update(updateData IUpdate) (*models.Role, error) {
	var role = models.Role{ID: updateData.RoleID}
	var updateMap = make(map[string]interface{})

	if updateData.Name == nil && updateData.Description == nil {
		return nil, errors.New("No fields to update")
	}

	if updateData.Name != nil {
		updateMap["Name"] = updateData.Name
	}

	if updateData.Description != nil {
		updateMap["Description"] = updateData.Description
	}

	result := r.getDB().Model(&role).Updates(updateMap)

	if result.Error != nil {
		return nil, result.Error
	}

	return &role, nil
}

func (r *Repository) Delete(deleteData IDelete) (*models.Role, error) {
	var role = models.Role{ID: deleteData.RoleID}

	verifyExistence := r.getDB().First(&role)

	if verifyExistence.Error != nil {
		return nil, verifyExistence.Error
	}

	result := r.getDB().Delete(&role)

	if result.Error != nil {
		return nil, result.Error
	}

	return &role, nil
}
