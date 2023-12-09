package taskproject

import (
	"gorm.io/gorm"
	"orb-api/models"
)


type(
	Repository struct{
		GetDB func() *gorm.DB
	}

	ICreate struct{
		TaskID		uint
		ProjectID	uint
	}

	IReadBy struct{
		ID			*uint
		TaskID		*uint
		ProjectID	*uint
		Limit		*int
	}

	IDelete struct{
		ID	uint
	}

	Interface interface{
		Create(ICreate)		(*models.TasksProject, error)
		ReadBy(IReadBy)		([]models.TasksProject, error)
		Delete(IDelete) 	(*models.TasksProject, error)
	}
)
