package roleservice

import (
	"orb-api/models"
	"orb-api/repositories/role"
)

type (
	RoleService struct {
		RoleRepo *role.Repository
	}

ICreateRole struct{
	Name string 
	Descripton string
}

IUpdateName struct{
	ID uint
	Name string
}

IUpdateDescription struct{
	ID uint
	Description string
}



Interface interface {
	CreateNewRole(ICreateRole) (*models.Role, error)
	UpdateRoleName(id uint, name string) (*models.Role, error)
	UpdateRoleDescription(id uint, description string) (*models.Role, error)
}

)