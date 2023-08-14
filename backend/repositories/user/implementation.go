package user

import (
	"errors"
	"gorm.io/gorm"
	"orb-api/models" //pacote que tem a declaração dos modelos (por isso models.User)
	"time"
)

func NewUserRepository(db *gorm.DB) Repository {
	return Repository{
		GetDB: func() *gorm.DB {
			return db
		},
	}
}

func (r *Repository) Create(create ICreate) error {
	var createError error
	var user models.User

	if len(create.Email) < emailMinLen || len(create.Email) > emailMaxLen {
		createError = errors.New("invalid email value")
		return createError
	}
	user.Email = create.Email

	if len(create.Name) < nameMinLen || len(create.Name) > nameMaxLen {
		createError = errors.New("invalid name value")
		return createError
	}
	user.Name = create.Name

	switch create.Status {
	case models.ActiveStatus:
		user.Status = models.ActiveStatus
	case models.ProcessingStatus:
		user.Status = models.ProcessingStatus
	case models.DisabledStatus:
		user.Status = models.DisabledStatus
	default:
		createError = errors.New("invalid status value")
		return createError
	}

	user.UpdatedAt = time.Now()

	if len(create.Password) < passwordMinLen {
		createError = errors.New("invalid password value")
		return createError
	}
	user.Password = create.Password

	result := r.GetDB().Create(&user)

	if result.Error != nil {
		createError = result.Error
		return createError
	}

	return nil
}

func (r *Repository) ReadAll(all IReadAll) ([]models.User, error) {
	var readError error
	var result *gorm.DB
	var userArray []models.User

	if all.Limit != nil {
		result = r.GetDB().Find(&userArray)
	} else {
		result = r.GetDB().Find(&userArray).Limit(*all.Limit)
	}

	if result.Error != nil {
		readError = result.Error
		return nil, readError
	}

	return userArray, nil
}

func (r *Repository) ReadBy(by IReadBy) ([]models.User, error) {
	var readError error
	var result *gorm.DB
	var fieldMap map[string]interface{}
	var userArray []models.User

	if by.Email != nil {
		email := *by.Email
		if len(email) < emailMinLen || len(email) > emailMaxLen {
			readError = errors.New("invalid email value")
			return nil, readError
		}

		fieldMap["Email"] = email
	}

	if by.Name != nil {
		name := *by.Name
		if len(name) < nameMinLen || len(name) > nameMaxLen {
			readError = errors.New("invalid name value")
			return nil, readError
		}

		fieldMap["Name"] = name
	}

	if by.UpdatedAt != nil {
		updatedAt := *by.UpdatedAt
		if updatedAt.After(time.Now()) {
			readError = errors.New("invalid updatedAt value")
			return nil, readError
		}

		fieldMap["updatedAt"] = updatedAt
	}

	switch status := *by.Status; status {
	case models.ActiveStatus:
		fieldMap["Status"] = models.ActiveStatus
	case models.ProcessingStatus:
		fieldMap["Status"] = models.ProcessingStatus
	case models.DisabledStatus:
		fieldMap["Status"] = models.DisabledStatus
	default:
		readError = errors.New("invalid status value")
		return nil, readError
	}

	if by.ID != nil {
		id := *by.ID
		fieldMap["ID"] = id
		userArray = make([]models.User, 1)
		result = r.GetDB().Where(fieldMap).First(&userArray[0])

		if result.Error != nil {
			readError = result.Error
			return nil, readError
		}

		return userArray, nil
	}

	if by.Limit != nil {
		limit := *by.Limit

		if limit < 1 {
			readError = errors.New("invalid limit value")
			return nil, readError
		}

		result = r.GetDB().Where(fieldMap).Find(&userArray).Limit(limit)
	} else {

		result = r.GetDB().Where(fieldMap).Find(&userArray)
	}

	if result.Error != nil {
		readError = result.Error
		return nil, readError
	}

	return userArray, readError
}

func (r *Repository) Update(update IUpdate) error {
	var updateError = errors.New("no fields to update")
	var result *gorm.DB
	var fieldMap map[string]interface{}
	var user models.User = models.User{ID: update.ID}

	if update.Email != nil {
		email := *update.Email
		if len(email) < emailMinLen || len(email) > emailMaxLen {
			updateError = errors.New("invalid email value")
			return updateError
		}

		fieldMap["Email"] = email
	}

	if update.Name != nil {
		name := *update.Name
		if len(name) < nameMinLen || len(name) > nameMaxLen {
			updateError = errors.New("invalid name value")
			return updateError
		}

		fieldMap["Name"] = name
	}

	user.UpdatedAt = time.Now()

	if update.Status != nil {
		var status uint = *update.Status

		switch status {
		case models.ActiveStatus:
			fieldMap["Status"] = models.ActiveStatus
		case models.ProcessingStatus:
			fieldMap["Status"] = models.ProcessingStatus
		case models.DisabledStatus:
			fieldMap["Status"] = models.DisabledStatus
		default:
			updateError = errors.New("invalid status value")
			return updateError
		}
	}

	if len(fieldMap) == 0 {
		result = r.GetDB().Where(user).Updates(fieldMap)

		if result.Error != nil {
			updateError = result.Error
			return updateError
		}

		return nil
	}

	return updateError
}

func (r *Repository) Delete(delete IDelete) error {
	var deleteError error
	var result *gorm.DB
	var user models.User = models.User{ID: delete.ID}

	verifyExistence := r.GetDB().First(&user)

	if verifyExistence.Error != nil {
		deleteError = verifyExistence.Error
		return deleteError
	}

	result = r.GetDB().Delete(&user)

	if result.Error != nil {
		deleteError = result.Error
		return deleteError
	}

	return nil
}
