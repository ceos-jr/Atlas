package task

import (
	"errors"
	"orb-api/models"
	"time"

	"gorm.io/gorm"
)

func NewTaskRepository(db *gorm.DB) Repository {
	return Repository{
		GetDB: func() *gorm.DB {
			return db
		},
	}
}

func ValidDeadline(deadline time.Time) bool {
	if deadline.Before(time.Now()) {
		return false
	}
	return true
}

func ValidTaskStatus(status uint) bool {
	_, valid := models.TaskStatus[status]

	return valid
}

func (r *Repository) ValidUser(id uint) bool {
	user := models.User{ID: id}

	verifyUser := r.GetDB().First(&user).Error

	if verifyUser != nil {
		return false
	}

	return true
}

func (r *Repository) ValidTask (id uint) bool {
	task := models.Task{ID: id}

	verifyTask := r.GetDB().First(&task).Error

	if verifyTask != nil{
		return false
	}

	return true
}

func (r *Repository) Create(createData ICreate) (*models.Task, error) {
	var task = models.Task{
		Description: createData.Description,
		CreatedBy:   createData.CreatedBy,
		AssignedTo:  createData.AssignedTo,
		Status:      createData.Status,
		Deadline:    createData.Deadline,
	}

	if !r.ValidUser(createData.CreatedBy) {
		return nil, errors.New("Invalid user passed to createBy")
	}

	if !r.ValidUser(createData.AssignedTo) {
		return nil, errors.New("Invalid user passed to assignedTo")
	}

	if !ValidTaskStatus(createData.Status) {
		return nil, errors.New("Invalid task status")
	}

	if !ValidDeadline(createData.Deadline) {
		return nil, errors.New("Invalid deadline")
	}

	result := r.GetDB().Create(&task)

	if result.Error != nil {
		return nil, result.Error
	}

	return &task, nil
}

func (r *Repository) ReadAll() ([]models.Task, error) {
	var taskArray []models.Task

	result := r.GetDB().Find(&taskArray)

	if result.Error != nil {
		return nil, result.Error
	}

	return taskArray, nil
}

func (r *Repository) ReadBy(readby IReadBy) ([]models.Task, error) {
	var fieldMap = make(map[string]interface{})
	var taskArray []models.Task
	var result *gorm.DB

	if readby.ID == nil &&
		readby.AssignedTo == nil &&
		readby.CreatedBy == nil &&
		readby.Status == nil &&
		readby.TimeRange == nil {
		return nil, errors.New("No fields to read")
	}

	if readby.ID != nil {
		fieldMap["id"] = *readby.ID
	}

	if readby.AssignedTo != nil {
		fieldMap["assigned_to"] = *readby.AssignedTo
	}

	if readby.CreatedBy != nil {
		fieldMap["created_by"] = *readby.CreatedBy
	}

	if readby.Status != nil {
		if !ValidTaskStatus(*readby.Status) {
			return nil, errors.New("Invalid task status")
		}

		fieldMap["status"] = *readby.Status
	}

	if readby.TimeRange != nil {
		result = r.GetDB().Where(fieldMap).Find(
			&taskArray, "deadline <= ?", *readby.TimeRange,
		)
	} else {
		result = r.GetDB().Where(fieldMap).Find(&taskArray)
	}

	if result.Error != nil {
		return nil, result.Error
	}

	return taskArray, nil
}

func (r *Repository) Update(updateData IUpdate) (*models.Task, error) {
	var task = models.Task{ID: updateData.ID}
	var fieldMap = make(map[string]interface{})

	if updateData.Description == nil &&
		updateData.AssignedTo == nil &&
		updateData.CreatedBy == nil &&
		updateData.Status == nil &&
		updateData.Deadline == nil {
		return nil, errors.New("No fields to update")
	}

	if !r.ValidTask(updateData.ID) {
		return nil, errors.New("Invalid task ID")
	}

	if updateData.Description != nil {
		fieldMap["description"] = *updateData.Description
	}

	if updateData.AssignedTo != nil {
		if !r.ValidUser(*updateData.AssignedTo) {
			return nil, errors.New("Invalid user passed to assignedTo")
		}

		fieldMap["assigned_to"] = *updateData.AssignedTo
	}

	if updateData.CreatedBy != nil {
		if !r.ValidUser(*updateData.CreatedBy) {
			return nil, errors.New("Invalid user passed to createBy")
		}

		fieldMap["created_by"] = *updateData.CreatedBy
	}

	if updateData.Status != nil {
		if !ValidTaskStatus(*updateData.Status) {
			return nil, errors.New("Invalid task status")
		}

		fieldMap["status"] = *updateData.Status
	}

	if updateData.Deadline != nil {
		if !ValidDeadline(*updateData.Deadline) {
			return nil, errors.New("Invalid deadline")
		}

		fieldMap["deadline"] = *updateData.Deadline
	}

	fieldMap["updated_at"] = time.Now()

	result := r.GetDB().Model(&task).Updates(fieldMap)

	if result.Error != nil {
		return nil, result.Error
	}

	return &task, nil
}

func (r *Repository) Delete(deleteData IDelete) (*models.Task, error) {
	var task = models.Task{ID: deleteData.ID}

	verifyExistence := r.GetDB().First(&task)

	if verifyExistence.Error != nil {
		return nil, verifyExistence.Error
	}

	result := r.GetDB().Delete(&task)

	if result.Error != nil {
		return nil, result.Error
	}

	return &task, nil
}
