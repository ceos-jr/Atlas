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