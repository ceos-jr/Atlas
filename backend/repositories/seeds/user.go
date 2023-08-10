package seeds

import (
	"github.com/bxcodec/faker/v4"
	"gorm.io/gorm"
	"orb-api/models"
)

func UserRandSeed(db *gorm.DB, size int) (*[]models.User, error) {
	var users = make([]models.User, size)

	for i := range users {
		users[i] = models.User{
			Name:     faker.Name(),
			Email:    faker.Email(),
			Password: faker.Password(),
			Status:   1,
		}
	}

	result := db.Create(users)

	if result.Error != nil {
		return nil, result.Error
	}

	return &users, nil
}
