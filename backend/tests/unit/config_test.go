package unit

import (
	"github.com/stretchr/testify/assert"
	"orb-api/config"
	"orb-api/models"
	"os"
	"testing"
)

var DBConfig config.DBConfig

func TestLoadEnv(test *testing.T) {
	assert := assert.New(test)

	loadEnvError := config.LoadEnv(".env")

	assert.Nil(loadEnvError, "Load env error should be nil")

	assert.Equal("test", os.Getenv("TEST"))
}

func TestCreateDBConnection(test *testing.T) {
	assert := assert.New(test)

	loadEnvError := config.LoadEnv(".env")

	assert.Nil(loadEnvError)

	DBConfig = config.DBConfig{
		Host:     os.Getenv("POSTGRES_HOST"),
		Port:     os.Getenv("POSTGRES_PORT"),
		DBName:   os.Getenv("POSTGRES_DB"),
		User:     os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		SLLMode:  os.Getenv("POSTGRES_SLL_MODE"),
	}

	connection, connectionError := config.CreateDBConnection(
		DBConfig,
	)

	assert.Nil(connectionError, "Connection error should be nil")

	if assert.NotNil(connection) {
		db, _ := connection.DB()
		defer db.Close()

		assert.Nil(db.Ping(), "Connection should be active")
	}
}

func TestMigrateDB(test *testing.T) {
	assert := assert.New(test)

	connection, _ := config.CreateDBConnection(
		DBConfig,
	)

	migrationError := config.MigrateDB(connection)

	assert.Nil(migrationError, "Migration error should be nil")

	if assert.NotNil(connection) {
		db, _ := connection.DB()
		hasTable := connection.Migrator().HasTable
		defer db.Close()

		assert.Equal(
			true, hasTable(&models.User{}), "Users should be in the database",
		)
		assert.Equal(
			true, hasTable(&models.Role{}), "Roles shoud be in the database",
		)
		assert.Equal(
			true, hasTable(&models.UserRole{}), "User roles should be in the database",
		)
		assert.Equal(
			true, hasTable(&models.Relation{}), "Relations should be in the database",
		)
		assert.Equal(
			true, hasTable(&models.Task{}), "Tasks should be in the database",
		)
		assert.Equal(
			true, hasTable(&models.Message{}), "Messages should be in the database",
		)
	}
}

func TestSetupDB(test *testing.T) {
	assert := assert.New(test)

	repository, setupError := config.SetupDB(".env")

	assert.Nil(setupError, "Setup error should be nil")

	if assert.NotNil(repository) {
		defer config.CloseDB(repository)

		assert.Nil(config.PingDB(repository), "Connection should be active")
	}
}
