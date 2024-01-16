package taskproject

import (
	"errors"
	"gorm.io/gorm"
	"orb-api/models"
)

func NewTaskProjectRepository(db *gorm.DB) Repository {
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

func (r *Repository) ValidTask(id uint) bool {
	task := models.Task{ID: id}

	verifyTask := r.GetDB().First(&task).Error

	if verifyTask != nil{
		return false
	}

	return true
}

func (r *Repository) Create(createData ICreate) (*models.TasksProject, error){

	var taskproject = models.TasksProject{
		ProjectID:	createData.ProjectID,
		TaskID:		createData.TaskID,
	}

	if !r.ValidProject(createData.ProjectID){
		return nil, errors.New("Invalid Project passed while creating taskproject")
	}

	if !r.ValidTask(createData.TaskID){
		return nil, errors.New("Invalid Task passed while creating taskproject")
	}

	result := r.GetDB().Create(&taskproject)

	if result.Error != nil{
		return nil, result.Error
	}

	return &taskproject, nil
}

func (r *Repository) ReadBy(readBy IReadBy) ([]models.TasksProject, error){
	var fieldMap = make(map[string]interface{})
	var taskprojectArray []models.TasksProject
	var result *gorm.DB

	if readBy.ID == nil &&
		readBy.ProjectID == nil &&
		readBy.TaskID == nil{
		return nil, errors.New("no fields to read")
	}

	if readBy.ID != nil {
		fieldMap["id"] = *readBy.ID
	}

	if readBy.ProjectID != nil {
		fieldMap["project_id"] = *readBy.ProjectID
	}

	if readBy.TaskID != nil {
		fieldMap["task_id"] = *readBy.TaskID
	}

	if readBy.Limit != nil {
		result = r.GetDB().Where(fieldMap).Find(&taskprojectArray).Limit(int(*readBy.Limit))
	} else {
		result = r.GetDB().Where(fieldMap).Find(&taskprojectArray)
	}

	if result.Error != nil {
		return nil, result.Error
	}

	return taskprojectArray, nil
}

func (r *Repository) Delete(deleteData IDelete) (*models.TasksProject, error) {
	var taskproject = models.TasksProject{ID: deleteData.ID}

	verifyExistence := r.GetDB().First(&taskproject)

	if verifyExistence.Error != nil {
		return nil, verifyExistence.Error
	}

	result := r.GetDB().Delete(&taskproject)

	if result.Error != nil {
		return nil, result.Error
	}

	return &taskproject, nil
}
