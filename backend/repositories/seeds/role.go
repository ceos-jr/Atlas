package seeds

import (
	repository "orb-api/repositories"
	"orb-api/repositories/role"

	"github.com/bxcodec/faker/v4"
)

func RoleRandSeed(repo *repository.Repository, size int) ([]role.ICreate, error) {
	var roles = make([]role.ICreate, size)

	for i := range roles {
		roles[i] = role.ICreate{
			Name:        faker.Name(),
			Description: faker.Sentence(),
		}

		result, _ := repo.Role.Create(roles[i])

		if result != nil {
			return nil, result
		}
	}

	return roles, nil
}
