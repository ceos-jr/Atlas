package project

import (
	"errors"
	"gorm.io/gorm"
	"orb-api/models"
)
func NewProjectRepository(db *gorm.DB) Repository {
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

func ValidProjectName(name string) bool{
	if len(name) < nameMinlen || len(name) > nameMaxlen{
		return false
	}
	return true
}

// Imaginando que um setor só possa ser identificado por números de 1 a 5
func ValidSector(numero uint) bool{
	return numero >= 1 && numero <= 5
}

func (r *Repository) Create(createData ICreate) (*models.Project, error){
	//Criar os models de sector para implantar sua validação de existencia
	//var sectorid = models.Sector{ID: createData.Sector}
	/*var adm = models.User{
		ID: createData.AdmID,
	}*/
	var project = models.Project{
		Name:		createData.Name,
		Sector:		createData.Sector,
		AdmID:		createData.AdmID,
	}


	if !ValidProjectName(createData.Name){
		return nil, errors.New("invalid name value")
	}
	//verifyAdmIDExistence := r.GetDB().First(&adm)

	/*if verifyAdmIDExistence.Error != nil{
		return nil, verifyAdmIDExistence.Error
	}
	//validar quando forem feitos os models de Adm e Sector

	/*verifyAdmIDExistence := r.GetDB().First(&adm)
 
	//Funções já criadas para validar a existencia do setor
	/*verifySectorIDExistence := r.GetDB().First(&sectorid)

	if verifySectorIDExistence.Error != nil {
		return nil, verifySectorIDExistence.Error
	}*/
	result := r.GetDB().Create(&project)

	if result.Error != nil{
		return nil, result.Error
	}
	
	return &project, nil
}
func (r *Repository) ReadBy(readBy IReadBy) ([]models.Project, error) {
	var fieldMap = make(map[string]interface{})
	var projectArray []models.Project
	var result *gorm.DB

	if readBy.ID == nil &&
		readBy.AdmID == nil &&
		readBy.Sector == nil &&
		readBy.Name == nil {
		return nil, errors.New("no fields to read")
	}

	if readBy.ID != nil {
		fieldMap["id"] = *readBy.ID
	}

	if readBy.Name != nil {
		if !ValidProjectName(*readBy.Name) {
			return nil, errors.New("invalid name")
		}

		fieldMap["name"] = *readBy.Name
	}

	if readBy.AdmID != nil {
		fieldMap["adm_id"] = *readBy.AdmID
	}

	if readBy.Sector != nil{
		fieldMap["sector"] = *readBy.Sector
	}

	if readBy.Limit != nil {
		result = r.GetDB().Where(fieldMap).Find(&projectArray).Limit(int(*readBy.Limit))
	} else {
		result = r.GetDB().Where(fieldMap).Find(&projectArray)
	}

	if result.Error != nil {
		return nil, result.Error
	}

	return projectArray, nil
}

func (r *Repository) Update(updateData IUpdate) (*models.Project, error) {
	var project = models.Project{ID: updateData.ID}
	var fieldMap = make(map[string]interface{})

	if updateData.Name == nil &&
		updateData.Sector == nil {
			return nil, errors.New("No fields to update")
		}
	
	if updateData.Name != nil {
		if !ValidProjectName(*updateData.Name){
			return nil, errors.New("Invalid project name")
		}
		fieldMap["name"] = *updateData.Name
	}

	if updateData.Sector != nil {
		if !ValidSector(*updateData.Sector){
			return nil, errors.New("Invalid Sector")
		}
		fieldMap["sector"] = *updateData.Sector
	}

	if updateData.AdmID != nil {
		fieldMap["adm_id"] = *updateData.AdmID
	}

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
                                                                                                                                                                                                                                                                                                       
