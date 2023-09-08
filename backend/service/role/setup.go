package role

import (
	"orb-api/models"
	"orb-api/repositories/role"
	"gorm.io/gorm"
)

// ServiceRole is a structure that represents the role service
type ServiceRole struct {
	roleRepository role.Repository
	db 		   *gorm.DB
}

// NewServiceRole is a function that returns a new instance of ServiceRole
func NewServiceRole(roleRepository role.Repository, db *gorm.DB) *ServiceRole {
	return &ServiceRole{
		roleRepository: roleRepository,
		db: db,
	}
}