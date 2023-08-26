package relation

import (
	"github.com/stretchr/testify/suite"
	"orb-api/models"
	"orb-api/repositories/relation"
	"testing"
)

const (
	UnmatchedLUserRoleID = "the left userRole does not match with expected"
	UnmatchedRUserRoleID = "the right userRole does not match with expected"
	UnmatchedStrongSide  = "the strong side does not match with expected"
)

func TestRelationRepository(test *testing.T) {
	suite.Run(test, new(TestSuite))
}

func (s *TestSuite) TestCreate() {
	auxRelations := make([]models.Relation, AuxRelationArraySize)

	for indexOut := range auxRelations {
		iCreate := s.MakeICreate()

		result, createError := s.Repo.Relation.Create(iCreate)

		s.Nil(createError)
		s.Equal(iCreate.LeftUserRoleID, result.LUserRoleID, UnmatchedLUserRoleID)
		s.Equal(iCreate.RightUserRoleID, result.RUserRoleID, UnmatchedRUserRoleID)
		s.NotNil(models.StrongSide.GetCode(iCreate.StrongSide))
		s.Equal(
			*models.StrongSide.GetCode(iCreate.StrongSide),
			result.StrongSide,
			UnmatchedStrongSide)

		auxRelations[indexOut] = *result
	}

	s.Nil(s.ClearAuxRelationsArray(auxRelations))
}

func (s *TestSuite) TestReadAll() {
	iReadAll := s.MakeIRealAll()
	mapMemory := make(map[uint]models.Relation)

	for _, rel := range s.MockRelation {
		s.NotEmpty(rel)
		mapMemory[rel.ID] = rel
	}

	result, readError := s.Repo.Relation.ReadAll(iReadAll)
	s.Nil(readError)

	if iReadAll.Limit != nil {
		s.Equal(int(*iReadAll.Limit), len(result))
	}

	for _, relDB := range result {
		relMemory, ok := mapMemory[relDB.ID]

		if ok {
			s.Equal(relMemory.ID, relDB.ID)
			s.Equal(relMemory.RUserRoleID, relDB.RUserRoleID, UnmatchedRUserRoleID)
			s.Equal(relMemory.LUserRoleID, relDB.LUserRoleID, UnmatchedLUserRoleID)
			s.Equal(relMemory.StrongSide, relDB.StrongSide, UnmatchedStrongSide)
		}
	}
}

func (s *TestSuite) TestReadBy() {
	var iReadBy = s.MakeIReadBy()
	var expected = make(map[uint]models.Relation)

	for _, rel := range s.MockRelation {
		id := (iReadBy.ID != nil) &&
			(*iReadBy.ID == rel.ID)
		rID := (iReadBy.RightUserRoleID != nil) &&
			(*iReadBy.RightUserRoleID == rel.RUserRoleID)
		lID := (iReadBy.LeftUserRoleID != nil) &&
			(*iReadBy.LeftUserRoleID == rel.LUserRoleID)
		ss := (iReadBy.StrongSide != nil) &&
			(*models.StrongSide.GetCode(*iReadBy.StrongSide) == rel.StrongSide)

		if id && rID && lID && ss {
			expected[rel.ID] = rel
		}
	}

	result, readError := s.Repo.Relation.ReadBy(iReadBy)

	s.Nil(readError)

	for _, rel := range result {
		expRel, ok := expected[rel.ID]

		s.True(ok, expRel)
		s.Equal(expRel.LUserRoleID, rel.LUserRoleID, UnmatchedLUserRoleID)
		s.Equal(expRel.RUserRole, rel.RUserRole, UnmatchedRUserRoleID)
		s.Equal(expRel.StrongSide, rel.StrongSide, UnmatchedStrongSide)
	}

	var expectedLen = len(expected)
	if iReadBy.Limit != nil {
		limit := int(*iReadBy.Limit)
		if limit > len(expected) {
			expectedLen = limit
		}
	}

	s.Equal(expectedLen, len(result))
}

func (s *TestSuite) TestUpdate() {
	var auxID uint
	var strongSide uint

	for _, rel := range s.MockRelation {
		auxID = rel.ID
		iUpdate := s.MakeIUpdate(&auxID)

		result, updateError := s.Repo.Relation.Update(iUpdate)

		s.Nil(updateError)

		updatedRel := *result
		auxID = iUpdate.ID
		s.Equal(updatedRel.ID, auxID)

		if iUpdate.RightUserRoleID != nil {
			auxID = *iUpdate.RightUserRoleID
		} else {
			auxID = rel.RUserRoleID
		}
		s.Equal(updatedRel.RUserRoleID, auxID, UnmatchedRUserRoleID)

		if iUpdate.LeftUserRoleID != nil {
			auxID = *iUpdate.LeftUserRoleID
		} else {
			auxID = rel.LUserRoleID
		}
		s.Equal(updatedRel.LUserRoleID, auxID, UnmatchedLUserRoleID)

		if iUpdate.StrongSide != nil {
			strongSide = *models.StrongSide.GetCode(*iUpdate.StrongSide)
		} else {
			strongSide = rel.StrongSide
		}
		s.Equal(updatedRel.StrongSide, strongSide, UnmatchedStrongSide)

		_, updateError = s.Repo.Relation.Update(relation.IUpdate{
			ID:              rel.ID,
			StrongSide:      models.StrongSide.GetName(rel.StrongSide),
			RightUserRoleID: &rel.RUserRoleID,
			LeftUserRoleID:  &rel.LUserRoleID,
		})

		s.Nil(updateError)
	}
}

func (s *TestSuite) TestDelete() {
	auxRelations, makeError := s.makeAuxRelationArray()

	s.Nil(makeError)

	for _, rel := range auxRelations {
		iDelete := relation.IDelete{ID: rel.ID}

		_, deleteError := s.Repo.Relation.Delete(iDelete)

		s.Nil(deleteError)
	}
}
