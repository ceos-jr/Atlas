package task

import (
	"errors"
	"gorm.io/gorm"
	"orb-api/models"
	"time"
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

func (r *Repository) Create(createData ICreate) (*models.Task, error) {
	var createdBy = models.User{ID: createData.CreatedBy}
	var assignedTo = models.User{ID: createData.AssignedTo}
	var task = models.Task{
		Description: createData.Description,
		CreatedBy:   createData.CreatedBy,
		AssignedTo:  createData.AssignedTo,
		Status:      createData.Status,
		Deadline:    createData.Deadline,
	}

	verifyCreateBy := r.GetDB().First(&createdBy)

	if verifyCreateBy.Error != nil {
		return nil, verifyCreateBy.Error
	}

	verifyAssignedTo := r.GetDB().First(&assignedTo)

	if verifyAssignedTo.Error != nil {
		return nil, verifyAssignedTo.Error
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
	var createdBy = models.User{ID: *updateData.CreatedBy}
	var assignedTo = models.User{ID: *updateData.AssignedTo}
	var task = models.Task{ID: updateData.ID}
	var fieldMap = make(map[string]interface{})

	if updateData.Description == nil &&
		updateData.AssignedTo == nil &&
		updateData.CreatedBy == nil &&
		updateData.Status == nil &&
		updateData.Deadline == nil {
		return nil, errors.New("No fields to update")
	}

	if updateData.Description != nil {
		fieldMap["description"] = *updateData.Description
	}

	if updateData.AssignedTo != nil {
		fieldMap["assigned_to"] = *updateData.AssignedTo
	}

	if updateData.CreatedBy != nil {
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

	verifyCreateBy := r.GetDB().First(&createdBy)

	if verifyCreateBy.Error != nil {
		return nil, verifyCreateBy.Error
	}

	verifyAssignedTo := r.GetDB().First(&assignedTo)

	if verifyAssignedTo.Error != nil {
		return nil, verifyAssignedTo.Error
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
