package userservicetest

import (
	userservice "orb-api/services/user"
	"testing"

	"github.com/stretchr/testify/suite"
)


func (suite *UserServiceTestSuit) TestUpdateEmail(){
	emailUser := "machomp@email.com"
	updateEmail, updateErr := userservice.Interface.UpdateEmail(uint(3), emailUser)

	suite.Nil(updateErr, "Update erros must be nil")
	suite.Equal(updateEmail, emailUser, "Expected to have the same email")

}



func (suite *UserServiceTestSuit) TestUpdateStatus() {
	status := uint(3)

	updateStatus, updateErr := userservice.UpdateStatus(uint(3), status)

	suite.Nil(updateErr, "Update erros must be nil")
	suite.Equal(updateStatus, status, "Expected to have the same email")
}



func TestUserRepository(t *testing.T) {
	suite.Run(t, new(UserServiceTestSuit))
}
