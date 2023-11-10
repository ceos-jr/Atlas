package project

import (
	//"orb-api/models"
	"orb-api/repositories/project"
)

type (
	Service struct {
		ProjectRepo* project.Repository
	}
)

type ICreate struct {
	Name		string
	//Sector		uint
	AdmID		uint
	EmployeeIDs []uint
}

//I did that to include EmployeesIDs 