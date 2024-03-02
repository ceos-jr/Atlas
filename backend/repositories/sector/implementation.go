package sector

import (
	"orb-api/models"
	"errors"
	"gorm.io/gorm"	
)

func NewSectorRepository(db *gorm.DB) Repository {
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

func (r *Repository) ValidUser(id uint) bool {
	user := models.User{ID: id}

	verifyUser := r.GetDB().First(&user).Error

	if verifyUser != nil {
		return false
	}

	return true
}

func (r *Repository) ValidSector(id uint) bool {
	sector := models.Sector{ID: id}

	verifySector := r.GetDB().First(&sector).Error

	if verifySector != nil{
		return false
	}

	return true
}

func ValidUserName(name string) bool {
	if len(name) < nameMinLen || len(name) > nameMaxLen {
		return false
	}
	return true
}

func ValidSectorName(name string) bool {
	if len(name) < nameMinLen || len(name) > nameMaxLen {
		return false
	}
	return true
}

func (r * Repository) Create(createData ICreate) (*models.Sector, error){
	var sector = models.Sector{
		Name: createData.Name,
		Description: createData.Description,
		AdmID: createData.AdmID,
	}

	if !r.ValidUser(createData.AdmID){
		return nil, errors.New("Invalid user passed to AdmID")
	}

	result := r.GetDB().Create(&sector)

	if result.Error != nil {
		return nil, result.Error
	}

	return &sector, nil

}

func (r *Repository) ReadAll() ([]models.Sector, error){
	var sectorsArray []models.Sector

	result := r.GetDB().Find(&sectorsArray)

	if result.Error != nil {
		return nil, result.Error
	}

	return sectorsArray, nil
}


func (r *Repository) ReadBy(readBy IReadBy) ([]models.Sector, error){
	var fieldMap = make(map[string]interface{})
	var sectorsArray []models.Sector
	var result *gorm.DB

	if readBy.ID == nil &&
		readBy.Name == nil &&
		readBy.AdmID == nil {
		return nil, errors.New("No fields to read")
	}

	if readBy.ID != nil {
		fieldMap["id"] = *readBy.ID
	}

	if readBy.Name != nil {
		if !ValidUserName(*readBy.Name) {
			return nil, errors.New("Invalid name")
		}

		fieldMap["name"] = *readBy.Name
	}

	if readBy.AdmID != nil {
		if !r.ValidUser(*readBy.AdmID) {
			return nil, errors.New("Invalid user passed to AdmID")
		}

		fieldMap["adm_id"] = *readBy.AdmID
	}

	if readBy.Limit != nil {
		result = r.GetDB().Where(fieldMap).Find(&sectorsArray).Limit(*readBy.Limit)
	} else {
		result = r.GetDB().Where(fieldMap).Find(&sectorsArray)
	}

	if result.Error != nil {
		return nil, result.Error
	}

	return sectorsArray, nil

}

func (r *Repository) Update(updateData IUpdate) (*models.Sector, error){
	var sector = models.Sector{ID: updateData.ID}
	var fieldMap = make(map[string]interface{})

	if updateData.Description == nil &&
		updateData.Name == nil &&
		updateData.AdmID == nil &&
		updateData.Members == nil &&
		updateData.Projects == nil {
		return nil, errors.New("No fields to update")
	}

	if !r.ValidSector(updateData.ID) {
		return nil, errors.New("Invalid sector ID")
	}

	if updateData.Description != nil {
		fieldMap["description"] = *updateData.Description
	}

	if updateData.Name != nil {
		if !ValidUserName(*updateData.Name) {
			return nil, errors.New("Invalid name")
		}

		fieldMap["name"] = *updateData.Name
	}

	if updateData.AdmID != nil {
		if !r.ValidUser(*updateData.AdmID) {
			return nil, errors.New("Invalid user passed to AdmID")
		}

		fieldMap["adm_id"] = *updateData.AdmID
	}

	if updateData.Members != nil {
		for _, member := range *updateData.Members {
			if !r.ValidUser(member) {
				return nil, errors.New("Invalid user passed to members")
			}
		}

		fieldMap["members"] = *updateData.Members
	}

	if updateData.Projects != nil {
		for _, project := range *updateData.Projects {
			if !r.ValidProject(project) {
				return nil, errors.New("Invalid project passed to projects")
			}
		}

		fieldMap["projects"] = *updateData.Projects
	}

	result := r.GetDB().Model(&sector).Updates(fieldMap)

	if result.Error != nil {
		return nil, result.Error
	}

	return &sector, nil
}

func (r *Repository) Delete(deleteData IDelete) (*models.Sector, error){
	var sector = models.Sector{ID: deleteData.ID}

	if !r.ValidSector(deleteData.ID) {
		return nil, errors.New("Invalid sector ID")
	}

	result := r.GetDB().Delete(&sector)

	if result.Error != nil {
		return nil, result.Error
	}

	return &sector, nil
}