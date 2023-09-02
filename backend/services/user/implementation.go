package userservice

import (
	"errors"
	"orb-api/repositories/user"
)

func UpdateEmail(email string) error{
	
	if !user.ValidUserEmail(email){
		return errors.New("Invalid email")
	}

	// if: email existente?


	return nil
}

func UpdateStatus(status *uint){
	// if status é valido:
	// se status = same status

}