package seeds

import (
	"github.com/bxcodec/faker/v4"
	"math/rand"
	repository "orb-api/repositories"
	"orb-api/repositories/user"
)

func UserRandSeed(repo *repository.Repository, size int) ([]user.ICreate, error) {
	var users = make([]user.ICreate, size)

	for i := range users {
		users[i] = user.ICreate{
			Name:     faker.Name(),
			Email:    faker.Email(),
			Status:   uint(rand.Intn(3) + 1),
			Password: faker.Password(),
		}

		result := repo.User.Create(users[i])

		if result != nil {
			return nil, result
		}
	}

	return users, nil
}
