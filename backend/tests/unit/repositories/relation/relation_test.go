package relation

import (
	"github.com/stretchr/testify/suite"
	"orb-api/models"
	"orb-api/repositories/relation"
	"testing"
)

const (
	UnmatchedLUserRoleId = "the left userRole does not match with expected"
	UnmatchedRUserRoleId = "the right userRole does not match with expected"
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
		s.Equal(iCreate.LeftUserRoleId, result.LUserRoleId, UnmatchedLUserRoleId)
		s.Equal(iCreate.RightUserRoleId, result.RUserRoleId, UnmatchedRUserRoleId)
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
			s.Equal(relMemory.RUserRoleId, relDB.RUserRoleId, UnmatchedRUserRoleId)
			s.Equal(relMemory.LUserRoleId, relDB.LUserRoleId, UnmatchedLUserRoleId)
			s.Equal(relMemory.StrongSide, relDB.StrongSide, UnmatchedStrongSide)
		}
	}
}

func (s *TestSuite) TestReadBy() {
	var iReadBy = s.MakeIReadBy()
	var expected = make(map[uint]models.Relation)

	for _, rel := range s.MockRelation {
		id := (iReadBy.ID != nil) && (*iReadBy.ID == rel.ID)
		rId := (iReadBy.RightUserRoleId != nil) && (*iReadBy.RightUserRoleId == rel.RUserRoleId)
		lId := (iReadBy.LeftUserRoleId != nil) && (*iReadBy.LeftUserRoleId == rel.LUserRoleId)
		ss := (iReadBy.StrongSide != nil) && (*models.StrongSide.GetCode(*iReadBy.StrongSide) == rel.StrongSide)

		if id && rId && lId && ss {
			expected[rel.ID] = rel
		}
	}

	result, readError := s.Repo.Relation.ReadBy(iReadBy)

	s.Nil(readError)

	for _, rel := range result {
		expRel, ok := expected[rel.ID]

		s.True(ok, expRel)
		s.Equal(expRel.LUserRoleId, rel.LUserRoleId, UnmatchedLUserRoleId)
		s.Equal(expRel.RUserRole, rel.RUserRole, UnmatchedRUserRoleId)
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
	var auxId uint
	var strongSide uint

	for _, rel := range s.MockRelation {
		auxId = rel.ID
		iUpdate := s.MakeIUpdate(&auxId)

		result, updateError := s.Repo.Relation.Update(iUpdate)

		s.Nil(updateError)

		updatedRel := *result
		auxId = iUpdate.ID
		s.Equal(updatedRel.ID, auxId)

		if iUpdate.RightUserRoleId != nil {
			auxId = *iUpdate.RightUserRoleId
		} else {
			auxId = rel.RUserRoleId
		}
		s.Equal(updatedRel.RUserRoleId, auxId, UnmatchedRUserRoleId)

		if iUpdate.LeftUserRoleId != nil {
			auxId = *iUpdate.LeftUserRoleId
		} else {
			auxId = rel.LUserRoleId
		}
		s.Equal(updatedRel.LUserRoleId, auxId, UnmatchedLUserRoleId)

		if iUpdate.StrongSide != nil {
			strongSide = *models.StrongSide.GetCode(*iUpdate.StrongSide)
		} else {
			strongSide = rel.StrongSide
		}
		s.Equal(updatedRel.StrongSide, strongSide, UnmatchedStrongSide)

		_, updateError = s.Repo.Relation.Update(relation.IUpdate{
			ID:              rel.ID,
			StrongSide:      models.StrongSide.GetName(rel.StrongSide),
			RightUserRoleId: &rel.RUserRoleId,
			LeftUserRoleId:  &rel.LUserRoleId,
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
