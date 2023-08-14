package seeds

import (
	"github.com/bxcodec/faker/v4"
	repository "orb-api/repositories"
	"orb-api/repositories/role"
)

func RoleRandSeed(repo *repository.Repository, size int) ([]role.ICreate, error) {
	var roles = make([]role.ICreate, size)

	for i := range roles {
		roles[i] = role.ICreate{
			Name:        faker.Name(),
			Description: faker.Sentence(),
		}

		result := repo.Role.Create(roles[i])

		if result != nil {
			return nil, result
		}
	}

	return roles, nil
}
