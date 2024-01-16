package userproject

import (
	"errors"
	"gorm.io/gorm"
	"orb-api/models"
)

func NewUserProjectRepository(db *gorm.DB) Repository {
	return Repository{
		GetDB: func() *gorm.DB {
			return db
		},
	}
}

func (r *Repository) ValidProject(id uint) bool {
	project := models.Project{ID: id}

	verifyProject := r.GetDB().First(&project).Error

	if verifyProject != nil {
		return false
	}

	return true
}

func (r *Repository) ValidUser(id uint) bool {
	user := models.User{ID: id}

	verifyUser := r.GetDB().First(&user).Error

	if verifyUser != nil {
		return false
	}

	return true
}

func (r *Repository) Create(createData ICreate) (*models.UsersProject, error){

	var userproject = models.UsersProject{
		ProjectID:	createData.ProjectID,
		UserID:		createData.UserID,
	}

	if !r.ValidProject(createData.ProjectID){
		return nil, errors.New("Invalid Project passed while creating userproject")
	}

	if !r.ValidUser(createData.UserID){
		return nil, errors.New("Invalid User passed while creating userproject")
	}

	result := r.GetDB().Create(&userproject)

	if result.Error != nil{
		return nil, result.Error
	}

	return &userproject, nil
}

func (r *Repository) ReadBy(readBy IReadBy) ([]models.UsersProject, error){
	var fieldMap = make(map[string]interface{})
	var userprojectArray []models.UsersProject
	var result *gorm.DB

	if readBy.ID == nil &&
		readBy.ProjectID == nil &&
		readBy.UserID == nil{
		return nil, errors.New("no fields to read")
	}

	if readBy.ID != nil {
		fieldMap["id"] = *readBy.ID
	}

	if readBy.ProjectID != nil {
		fieldMap["project_id"] = *readBy.ProjectID
	}

	if readBy.UserID != nil {
		fieldMap["user_id"] = *readBy.UserID
	}

	if readBy.Limit != nil {
		result = r.GetDB().Where(fieldMap).Find(&userprojectArray).Limit(int(*readBy.Limit))
	} else {
		result = r.GetDB().Where(fieldMap).Find(&userprojectArray)
	}

	if result.Error != nil {
		return nil, result.Error
	}

	return userprojectArray, nil
}

func (r *Repository) Delete(deleteData IDelete) (*models.UsersProject, error) {
	var userproject = models.UsersProject{ID: deleteData.ID}

	verifyExistence := r.GetDB().First(&userproject)

	if verifyExistence.Error != nil {
		return nil, verifyExistence.Error
	}

	result := r.GetDB().Delete(&userproject)

	if result.Error != nil {
		return nil, result.Error
	}

	return &userproject, nil
}
