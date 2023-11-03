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


///////////////////////////////////////////////////////////////////
//  var updateData do tipo IUpdate
//  
func (r *Repository) Update(updateData IUpdate) (*models.Project, error) {
	var project = models.Project{ID: updateData.ID}
	var fieldMap = make(map[string]interface{})

	if updateData.Name == nil &&
		updateData.Sector == nil {
			return nil, errors.New("No fields to update")
		}
	
	if updateData.Name != nil {
		if !r.ValidProjectName(*&updateData.Name) {
			return nil, errors.New("Invalid Project Name")
		}
		fieldMap["project_name"] = *&updateData.Name
	}

	if updateData.Sector != nil {
		fieldMap["sector_number"] = *&updateData.Sector
	}

	fieldMap["updated_at"] = time.Now()
	result := r.GetDB().Model(&project).Updates(fieldMap)

	if result.Error != nil {
		return nil, result.Error
	}

	return &project, nil
}

func (r *Repository) Delete(deleteData IDelete) (*models.Project, error) {
	var project = models.Project{ID: deleteData.ID}
	verifyExistence := r.GetDB().First(&project)

	if verifyExistence.Error != nil {
		return nil, verifyExistence.Error
	}

	result := r.GetDB().Delete(&project)

	if result.Error != nil {
		return nil, result.Error
	}

	return &project, nil
}


// Criar uma função pra validar o número do setor 
                                                                                                                                                                                                                                                                                                       