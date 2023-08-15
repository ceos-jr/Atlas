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

func (r *Repository) Create(createData ICreate) error {
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
		return verifyCreateBy.Error
	}

	verifyAssignedTo := r.GetDB().First(&assignedTo)

	if verifyAssignedTo.Error != nil {
		return verifyAssignedTo.Error
	}

	if !ValidTaskStatus(createData.Status) {
		return errors.New("Invalid task status")
	}

	if !ValidDeadline(createData.Deadline) {
		return errors.New("Invalid deadline")
	}

	result := r.GetDB().Create(&task)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
