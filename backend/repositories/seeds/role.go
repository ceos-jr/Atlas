package seeds

import (
	"github.com/bxcodec/faker/v4"
	"gorm.io/gorm"
	"orb-api/models"
)

func RoleRandSeed(db *gorm.DB, size int) (*[]models.Role, error) {
	var roles = make([]models.Role, size)

	for i, _ := range roles {
		roles[i] = models.Role{
			Name:        faker.Name(),
			Description: faker.Sentence(),
		}
	}

	result := db.Create(roles)

	if result.Error != nil {
		return nil, result.Error
	}

	return &roles, nil
}
