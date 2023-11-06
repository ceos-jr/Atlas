package project

import (
	"errors"
	"gorm.io/gorm"
	"orb-api/models"
	"time"
)
func NewProjectRepository(db *gorm.DB) Repository {
	return Repository{
		GetDB: func() *gorm.DB {
			return db
		},
	}
}


func ValidProjectName(name string) bool{
	if len(name) < nameMinlen || len(name) > nameMaxlen{
		return false
	}
	return true
}

func (r *Repository) Create(createData ICreate) (*models.Project, error){
	//Criar os models de sector para implantar sua validação de existencia
	//var sectorid = models.Sector{ID: createData.Sector}
	var adm = models.User{ID: createData.AdmID}
	var project = models.Project{
		Name:		createData.Name,
		Sector:		createData.Sector,
		AdmID:		createData.AdmID,
	}


	if !ValidProjectName(createData.Name){
		return nil, errors.New("invalid name value")
	}
	verifyAdmIDExistence := r.GetDB().First(&adm)

	if verifyAdmIDExistence.Error != nil{
		return nil, verifyAdmIDExistence.Error
	}
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
		fieldMap["admid"] = *readBy.AdmID
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


///////////////////////////////////////////////////////////////////
//  var updateData do tipo IUpdate
//  
func (r *Repository) Update(updateData IUpdate) (*models.Project, error) {
	var project = models.Project{ID: updateData.ID}
	var fieldMap = make(map[string]interface{})

	if updateData.Name == nil &&
		updateData.Sector == nil {
			return nil, errors.New("no fields to update")
		}
	
	if updateData.Name != nil {
		//if !r.ValidProjectName(*&updateData.Name) {
		//	return nil, errors.New("Invalid Project Name")
		//}
		fieldMap["project_name"] = *updateData.Name
	}

	if updateData.Sector != nil {
		fieldMap["sector_number"] = *updateData.Sector
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
                                                                                                                                                                                                                                                                                                       