package userservicerepotest

import (
	"testing"

	"github.com/stretchr/testify/suite"
)


func (suite *UserServiceRepoTestSuit) TestUpdateEmail(){
	emailUser := "machomp@email.com"
	updateEmail, updateErr := suite.service.user.UpdateEmail(uint(3),&emailUser)

	suite.Nil(updateErr, "Update erros must be nil")
	suite.Equal(updateEmail, emailUser, "Expected to have the same email")

}



func (suite *UserServiceRepoTestSuit) TestUpdateStatus() {
	status := uint(3)

	updateStatus, updateErr := suite.service.user.UpdateStatus(uint(3), status)

	suite.Nil(updateErr, "Update erros must be nil")
	suite.Equal(updateStatus, status, "Expected to have the same email")
}



func TestUserRepository(t *testing.T) {
	suite.Run(t, new(UserRepoTestSuite))
}
