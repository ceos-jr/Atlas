package project

import (
	"gorm.io/gorm"
	"orb-api/models"
)
const(
	nameMaxlen = 128
	nameMinlen = 3
)
type(
	Repository struct{
		GetDB func() *gorm.DB
	}

	ICreate struct{
		Name		string
		Sector		uint
		AdmID		uint
	}

	IReadAll struct{
		Limit		*uint
	}

	IReadBy struct{
		ID			*uint
		Name		*string
		Sector		*uint
		AdmID		*uint
		Limit		*uint
	}
	IUpdate struct{
		ID 			uint
		Name 		*string
		Sector		*uint
	}

	IDelete struct{
		ID 			uint
	}

	Interface interface{
		Create(ICreate)		(*models.Project, error)
		ReadAll(IReadAll)	([]models.Project, error)
		ReadBy(IReadBy)		([]models.Project, error)
		Update(IUpdate)		(*models.Project, error)
		Delete(IDelete)		(*models.Project, error)
	}
)

// IReadByName, IReadBySectorID, IReadByAdmID