package userproject

import (
	"gorm.io/gorm"
	"orb-api/models"
)

type(
	Repository struct{
		GetDB func() *gorm.DB
	}

	ICreate struct{
		UserID		uint
		ProjectID	uint
	}

	IReadBy struct{
		ID			*uint
		UserID		*uint
		ProjectID	*uint
		Limit		*int
	}

	IDelete struct{
		ID	uint
	}

	Interface interface{
		Create(ICreate)		(*models.UsersProject, error)
		ReadBy(IReadBy)		([]models.UsersProject, error)
		Delete(IDelete) 	(*models.UsersProject, error)
	}
)
