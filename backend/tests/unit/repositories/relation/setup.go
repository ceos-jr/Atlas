package relation

import (
	"github.com/bxcodec/faker/v4"
	"github.com/stretchr/testify/suite"
	"math/rand"
	"orb-api/config"
	"orb-api/models"
	repository "orb-api/repositories"
	"orb-api/repositories/relation"
	"orb-api/repositories/role"
	"orb-api/repositories/user"
	userrole "orb-api/repositories/user_role"
)

const (
	MockArraySize        = 100
	AuxRelationArraySize = 100
)

type TestSuite struct {
	suite.Suite
	Repo         *repository.Repository
	MockUserRole []models.UserRole
	MockUser     []models.User
	MockRole     []models.Role
	MockRelation []models.Relation
}

func (s *TestSuite) SetupSuite() {
	repo, setupError := config.SetupDB("../.env")

	if setupError != nil {
		panic(setupError)
	}

	s.Repo = repo
	s.MockUserRole = make([]models.UserRole, MockArraySize)
	s.MockUser = make([]models.User, MockArraySize)
	s.MockRole = make([]models.Role, MockArraySize)
	s.MockRelation = make([]models.Relation, MockArraySize)

	s.SetupMocks()
}

func (s *TestSuite) SetupMocks() {
	for index := range s.MockUser {
		createdUser, createError := s.Repo.User.Create(user.ICreate{
			Name:     faker.Name(),
			Email:    faker.Email(),
			Status:   models.UStatusActive,
			Password: faker.Password(),
		})

		if createError != nil {
			panic(createError)
		}

		s.MockUser[index] = *createdUser
	}

	for index := range s.MockRole {
		createdRole, createError := s.Repo.Role.Create(role.ICreate{
			Name:        faker.Name(),
			Description: faker.Sentence(),
		})

		if createError != nil {
			panic(createError)
		}

		s.MockRole[index] = *createdRole
	}

	for index := range s.MockUserRole {
		userRole, createError := s.Repo.UserRole.Create(userrole.ICreate{
			RoleID: s.MockRole[index].ID,
			UserID: s.MockUser[index].ID,
		})

		if createError != nil {
			panic(createError)
		}

		s.MockUserRole[index] = *userRole
	}

	for index := range s.MockRelation {
		createdRelation, createErr := s.Repo.Relation.Create(s.MakeICreate())

		if createErr != nil {
			panic(createErr.Error())
		}

		s.MockRelation[index] = *createdRelation
	}
}

func (s *TestSuite) TearDownSuite() {
	var tearDownPanic error

	for _, rel := range s.MockRelation {
		_, tearDownPanic = s.Repo.Relation.Delete(
			relation.IDelete{ID: rel.ID})

		if tearDownPanic != nil {
			panic(tearDownPanic.Error())
		}
	}

	for _, userRole := range s.MockUserRole {
		_, tearDownPanic = s.Repo.UserRole.Delete(userrole.IDelete{
			UserRoleID: userRole.ID,
		})

		if tearDownPanic != nil {
			panic(tearDownPanic.Error())
		}
	}

	for _, userRef := range s.MockUser {
		_, tearDownPanic = s.Repo.User.Delete(user.IDelete{
			ID: userRef.ID,
		})

		if tearDownPanic != nil {
			panic(tearDownPanic.Error())
		}
	}

	for _, roleRef := range s.MockRole {
		_, tearDownPanic = s.Repo.Role.Delete(role.IDelete{
			RoleID: roleRef.ID,
		})

		if tearDownPanic != nil {
			panic(tearDownPanic.Error())
		}
	}
}

func (s *TestSuite) MakeIReadBy() relation.IReadBy {
	var iReadBy = relation.IReadBy{}
	var fieldRandomChoice = rand.Intn(2 ^ 5)

	codeID := 0b00001
	var id *uint = nil
	if fieldRandomChoice&codeID == codeID {
		auxID := s.MockUser[RandIndex(nil, len(s.MockUser))].ID
		id = &auxID
	}
	iReadBy.ID = id

	codeStrongSide := 0b00010
	var strongSide *string = nil
	if fieldRandomChoice&codeStrongSide == codeStrongSide {
		auxStrongSide := models.StrongSide.RandState().Name
		strongSide = &auxStrongSide
	}
	iReadBy.StrongSide = strongSide

	codeRUserRoleID := 0b00100
	var rUserRoleID *uint = nil
	if fieldRandomChoice&codeRUserRoleID == codeRUserRoleID {
		auxRUserRole := s.MockUserRole[RandIndex(nil, len(s.MockUserRole))].ID
		rUserRoleID = &auxRUserRole
	}
	iReadBy.RightUserRoleID = rUserRoleID

	codeLUserRoleID := 0b01000
	var lUserRoleID *uint = nil
	if fieldRandomChoice&codeLUserRoleID == codeLUserRoleID {
		auxRUserRole := s.MockUserRole[RandIndex(nil, len(s.MockUserRole))].ID
		lUserRoleID = &auxRUserRole
	}
	iReadBy.LeftUserRoleID = lUserRoleID

	codeLimit := 0b10000
	var limit *uint = nil
	if fieldRandomChoice&codeLimit == codeLimit {
		auxLimit := uint(rand.Intn(len(s.MockRelation)))
		limit = &auxLimit
	}
	iReadBy.Limit = limit

	codeEmpty := 0b00000
	if fieldRandomChoice&codeEmpty == codeEmpty {
		auxRelation := s.MockRelation[RandIndex(nil, len(s.MockRelation))]

		iReadBy.ID = &auxRelation.ID
		iReadBy.RightUserRoleID = &auxRelation.RUserRoleID
		iReadBy.LeftUserRoleID = &auxRelation.LUserRoleID
		iReadBy.StrongSide = models.StrongSide.GetName(auxRelation.StrongSide)
		iReadBy.Limit = nil
	}

	return iReadBy
}

func (s *TestSuite) MakeIRealAll() relation.IReadAll {
	if rand.Intn(2) == 0 {
		return relation.IReadAll{Limit: nil}
	}

	limit := uint(rand.Intn(AuxRelationArraySize) + 1)
	return relation.IReadAll{Limit: &limit}
}

func (s *TestSuite) MakeICreate() relation.ICreate {
	var iCreate relation.ICreate

	rightIndex := RandIndex(nil, len(s.MockUserRole))
	leftIndex := RandIndex(&rightIndex, len(s.MockUserRole))

	iCreate.RightUserRoleID = s.MockUserRole[rightIndex].ID
	iCreate.LeftUserRoleID = s.MockUserRole[leftIndex].ID
	iCreate.StrongSide = models.StrongSide.RandState().Name

	return iCreate
}

func RandIndex(diffFrom *int, arrayLen int) int {
	if diffFrom == nil {
		return rand.Intn(arrayLen)
	}

	diffIndex := *diffFrom
	virtualArrayLen := arrayLen - 1
	lastIndex := virtualArrayLen
	randIndex := rand.Intn(virtualArrayLen)

	if lastIndex == diffIndex {
		return randIndex
	}

	if randIndex == diffIndex {
		return lastIndex
	}

	return randIndex
}

func (s *TestSuite) MakeIUpdate(fromID *uint) relation.IUpdate {
	var iUpdate relation.IUpdate
	var fieldRandomChoice = rand.Intn(2 ^ 3)
	var randIndex *int = nil

	codeStrongSide := 0b001
	if fieldRandomChoice&codeStrongSide == codeStrongSide {
		stateName := models.StrongSide.RandState().Name
		iUpdate.StrongSide = &stateName
	}

	codeRightID := 0b010
	if fieldRandomChoice&codeRightID == codeRightID {
		auxRandIndex := RandIndex(nil, len(s.MockUserRole))
		randIndex = &auxRandIndex
		iUpdate.RightUserRoleID = &s.MockUserRole[auxRandIndex].ID
	}

	codeLeftID := 0b100
	if fieldRandomChoice&codeLeftID == codeLeftID {
		auxRandIndex := RandIndex(randIndex, len(s.MockUserRole))
		iUpdate.LeftUserRoleID = &s.MockUserRole[auxRandIndex].ID
	}

	codeEmpty := 0b000
	if fieldRandomChoice&codeEmpty == codeEmpty {
		stateName := models.StrongSide.RandState().Name
		iUpdate.StrongSide = &stateName

		auxRandIndex := RandIndex(nil, len(s.MockUserRole))
		randIndex = &auxRandIndex
		iUpdate.RightUserRoleID = &s.MockUserRole[auxRandIndex].ID

		auxRandIndex = RandIndex(randIndex, len(s.MockUserRole))
		iUpdate.LeftUserRoleID = &s.MockUserRole[auxRandIndex].ID
	}

	auxRandIndex := RandIndex(nil, len(s.MockRelation))

	if fromID != nil {
		iUpdate.ID = *fromID
	} else {
		iUpdate.ID = s.MockRelation[auxRandIndex].ID
	}

	return iUpdate
}

func (s *TestSuite) makeAuxRelationArray() ([]models.Relation, error) {
	var relations = make([]models.Relation, AuxRelationArraySize)

	for index := range relations {
		iCreate := s.MakeICreate()

		result, createError := s.Repo.Relation.Create(iCreate)

		if createError != nil {
			return nil, createError
		}

		relations[index] = *result
	}

	return relations, nil
}

func (s *TestSuite) ClearAuxRelationsArray(relations []models.Relation) error {
	for _, rel := range relations {
		_, clearError := s.Repo.Relation.Delete(
			relation.IDelete{ID: rel.ID})

		if clearError != nil {
			return clearError
		}
	}

	return nil
}
