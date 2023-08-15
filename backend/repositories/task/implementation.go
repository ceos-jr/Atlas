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

func (r *Repository) ReadAll() ([]models.Task, error) {
  var taskArray []models.Task

  result := r.GetDB().Find(&taskArray)

  if result.Error != nil {
    return nil, result.Error
  }

  return taskArray, nil
}

func (r *Repository) ReadBy(readby IReadBy) ([]models.Task, error) {
  var fieldMap map[string]interface{} 
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
    fieldMap["ID"] = *readby.ID
  }

  if readby.AssignedTo != nil {
    fieldMap["AssignedTo"] = *readby.AssignedTo
  }

  if readby.CreatedBy != nil {
    fieldMap["CreatedBy"] = *readby.CreatedBy
  }

  if readby.Status != nil {
    if !ValidTaskStatus(*readby.Status) {
      return nil, errors.New("Invalid task status")
    }

    fieldMap["AssignedTo"] = *readby.Status
  }

  if readby.TimeRange != nil {
    result = r.GetDB().Where(fieldMap, "deadline <= ?", *readby.TimeRange).Find(&taskArray)
  } else {
    result = r.GetDB().Where(fieldMap).Find(&taskArray)
  }

  if result.Error != nil {
    return nil, result.Error
  }

  return nil, nil
}

func (r *Repository) Update(updateData IUpdate) error {
 	var createdBy = models.User{ID: *updateData.CreatedBy}
	var assignedTo = models.User{ID: *updateData.AssignedTo}
  var task = models.Task{ID: updateData.ID}
  var fieldMap map[string]interface{}
 
  if updateData.Description == nil && 
    updateData.AssignedTo == nil &&
    updateData.CreatedBy == nil && 
    updateData.Status == nil && 
    updateData.Deadline == nil {
    return errors.New("No fields to read")
  }
  
  if updateData.Description != nil {
    fieldMap["Description"] = *updateData.Description
  }

  if updateData.AssignedTo != nil {
    fieldMap["AssignedTo"] = *updateData.AssignedTo
  } 

  if updateData.CreatedBy != nil {
    fieldMap["CreatedBy"] = *updateData.CreatedBy
  }

  if updateData.Status != nil {
    if !ValidTaskStatus(*updateData.Status) {
      return errors.New("Invalid task status")
    }

    fieldMap["Status"] = *updateData.Status
  }

  if updateData.Deadline != nil {
    if !ValidDeadline(*updateData.Deadline) {
      return errors.New("Invalid deadline")
    }

    fieldMap["Deadline"] = *updateData.Deadline
  }
  
  verifyCreateBy := r.GetDB().First(&createdBy)

	if verifyCreateBy.Error != nil {
		return verifyCreateBy.Error
	}

	verifyAssignedTo := r.GetDB().First(&assignedTo)

	if verifyAssignedTo.Error != nil {
		return verifyAssignedTo.Error
	}

  fieldMap["UpdatedAt"] = time.Now()

  result := r.GetDB().Model(&task).Updates(fieldMap)

  if result.Error != nil {
    return result.Error
  }

  return nil
}

func (r *Repository) Delete(deleteData IDelete) error {
	var task = models.Task{ID: deleteData.ID}

	verifyExistence := r.GetDB().First(&task)

	if verifyExistence.Error != nil {
		return verifyExistence.Error
	}

	result := r.GetDB().Delete(&task)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
