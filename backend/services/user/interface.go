package user


import (
	"orb-api/models"
	"orb-api/repositories/user"
	"orb-api/repositories/userproject"
	"orb-api/repositories/project"
)

const (
	emailMaxLen    = 128
	emailMinLen    = 3
	nameMaxLen     = 128
	nameMinLen     = 3

	passwordMinLen = 8
)

type (
	Service struct {
		UserRepo *user.Repository
		UserProjectRepo	*userproject.Repository
		ProjectRepo *project.Repository
	}

	Interface interface {
		NewUser(name, email, password string) (*models.User, error)

		UpdateName(id uint, name string) (*models.User, error)
		UpdateEmail(id uint, email string) (*models.User, error)
		UpdatePassword(id uint, password string) (*models.User, error)
		UpdateStatus(id uint, status uint) (*models.User, error)
		DeleteUser(id uint) (*models.User, error)
		SortProjects(id uint) ([]models.Project, error)
		ReadUser(read user.IReadBy) ([]models.User, error)
	}
)
