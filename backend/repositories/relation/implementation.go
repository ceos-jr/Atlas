package relation

import (
	"errors"
	"gorm.io/gorm"
	"orb-api/models"
)

func NewRepository(db *gorm.DB) Repository {
	return Repository{
		GetDB: func() *gorm.DB {
			return db
		},
	}
}

func (r *Repository) Create(create ICreate) (*models.Relation, error) {
	var createError error = nil
	var relation = models.Relation{}

	strongSide := models.StrongSide.GetCode(create.StrongSide)

	if strongSide == nil {
		createError = errors.New(CErrorInvalidStrongSide)
		return nil, createError
	}

	relation.StrongSide = *strongSide

	relation.LUserRoleID = create.LeftUserRoleID

	relation.RUserRoleID = create.RightUserRoleID

	result := r.GetDB().Model(&models.Relation{}).Create(&relation)

	if result.Error != nil {
		createError = result.Error
		return nil, createError
	}

	return &relation, nil
}

func (r *Repository) ReadBy(readBy IReadBy) ([]models.Relation, error) {
	var readError error = nil
	var result *gorm.DB
	var relations []models.Relation
	var refRelation = models.Relation{}

	if readBy.ID == nil && readBy.StrongSide == nil &&
		readBy.RightUserRoleID == nil && readBy.LeftUserRoleID == nil {
		readError = errors.New(RErrorEmptyReadBy)
		return nil, readError
	}

	if readBy.ID != nil {
		id := *readBy.ID

		refRelation.ID = id
	}

	if readBy.RightUserRoleID != nil {
		RUserRoleID := *readBy.RightUserRoleID

		refRelation.RUserRoleID = RUserRoleID
	}

	if readBy.LeftUserRoleID != nil {
		LUserRoleID := *readBy.LeftUserRoleID

		refRelation.LUserRoleID = LUserRoleID
	}

	if readBy.StrongSide != nil {
		strongSide := models.StrongSide.GetCode(*readBy.StrongSide)

		if strongSide == nil {
			readError = errors.New(RErrorInvalidStrongSide)
			return nil, readError
		}

		refRelation.StrongSide = *strongSide
	}

	if readBy.Limit == nil {
		result = r.GetDB().Model(&models.Relation{}).
			Where(refRelation).Find(&relations)
	} else {
		result = r.GetDB().Model(&models.Relation{}).
			Where(refRelation).Find(&relations).Limit(int(*readBy.Limit))
	}

	if result.Error != nil {
		readError = result.Error
		return nil, readError
	}

	return relations, nil
}

func (r *Repository) ReadAll(readAll IReadAll) ([]models.Relation, error) {
	var readError error = nil
	var relations []models.Relation

	readError = r.matchReadLimit(readAll.Limit, &relations)

	if readError != nil {
		return nil, readError
	}

	return relations, nil
}

func (r *Repository) Update(update IUpdate) (*models.Relation, error) {
	var updateError error = nil
	var relation = models.Relation{ID: update.ID}

	if update.StrongSide == nil &&
		update.LeftUserRoleID == nil &&
		update.RightUserRoleID == nil {
		updateError = errors.New(UErrorEmptyUpdate)
		return nil, updateError
	}

	if update.StrongSide != nil {
		strongSide := models.StrongSide.GetCode(*update.StrongSide)

		if strongSide == nil {
			updateError = errors.New(UErrorInvalidStrongSide)
			return nil, updateError
		}

		relation.StrongSide = *strongSide
	}

	if update.RightUserRoleID != nil {
		relation.RUserRoleID = *update.RightUserRoleID
	}

	if update.RightUserRoleID != nil {
		relation.RUserRoleID = *update.RightUserRoleID
	}

	if update.LeftUserRoleID != nil {
		relation.LUserRoleID = *update.LeftUserRoleID
	}

	result := r.GetDB().Updates(&relation)

	if result.Error != nil {
		updateError = result.Error
		return nil, updateError
	}

	return &relation, nil
}

func (r *Repository) Delete(delete IDelete) (*models.Relation, error) {
	var deleteError error
	var relation = models.Relation{}

	result := r.GetDB().Delete(&relation, delete.ID)

	if result.Error != nil {
		deleteError = result.Error
		return nil, deleteError
	}

	return &relation, nil
}

func (r *Repository) matchReadLimit(limit *uint, relations *[]models.Relation) error {
	var result *gorm.DB

	if limit == nil {
		result = r.GetDB().Model(models.Relation{}).Find(&relations)
	} else {
		result = r.GetDB().Model(models.Relation{}).Limit(int(*limit)).Find(&relations)
	}

	if result.Error != nil {
		return result.Error
	}

	return nil
}
