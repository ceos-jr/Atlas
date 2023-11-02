package project

import (
	"errors"
	"gorm.io/gorm"
	"orb-api/models"
)
func NewProjectRepository(db *gorm.DB) Repository {
	return Repository{
		GetDB: func() *gorm.DB{
			return db
		}
	}
}

func ValidProjectName(name string) bool{
	if len(name) < nameMinlen || len(name) > nameMaxlen{
		return false
	}
	return true
}
func ValidProjectAdmID(sector uint) bool{
	
}

func (r *Repository) Create(createData ICreate) (*models.Project, error){
	var project = models.Project{
		Name:		createData.Name,
		Sector:		createData.Sector,
		AdmID:		createData.AdmID,
	}
	if !ValidProjectName(createData.Name){
		return nil, errors.New("invalid name value")
	}
	if !ValidProjectSector(createData.Sector){
		return nil, errors.New("invalid sector value")
	}
	if !ValidProjectAdmID(createData.AdmID){
		return nil, errors.New("invalid Admid value")
	}
	
	result := r.GetDB().Create(&project)

	if result.Error != nil{
		return nil, result.Error
	}
	
	return &project, nil
}