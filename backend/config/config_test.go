package config

import (
	"errors"
	"orb-api/models"
  "orb-api/utils"
	"os"
	"reflect"
	"testing"
	"time"
)

// TestUserCRUD tests the CRUD operations for the User model.
func TestUserCRUD(t *testing.T) {
	ErrorLabel := "Test"
	if envError := LoadEnv(".env"); envError != nil {
		envError.AddLabel(ErrorLabel)
		t.Fatalf(envError.Error())
	}

	config := DBConfig{
		Host:     os.Getenv("POSTGRES_HOST"),
		Port:     os.Getenv("POSTGRES_PORT"),
		DBName:   os.Getenv("POSTGRES_DB"),
		User:     os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		SLLMode:  os.Getenv("POSTGRES_SLL_MODE"),
	}

	db, configError := CreateDBConnection(config)

	if configError != nil {
		configError.AddLabel(ErrorLabel)
		t.Fatal(configError.Error())
	}

	// Auto-migrate the User model
	migrateError := utils.NewError(ErrorLabel, db.AutoMigrate(&models.User{}))
	if migrateError != nil {
		t.Fatalf(migrateError.Error())
	}

	// Create a new user and save it in the database
	user := models.User{
		Name:      "John Doe",
		Email:     "john@example.com",
		Status:    1,
		Password:  "secretpassword",
		UpdatedAt: time.Now(),
	}

	creteError := db.Create(&user).Error
	if creteError != nil {
		creteError = utils.NewError(ErrorLabel, creteError)
		t.Fatalf(creteError.Error())
	}

	// Read the user from the database
	var retrievedUser models.User
	readError := db.First(&retrievedUser, user.ID).Error
	if readError != nil {
		readError = utils.NewError(ErrorLabel, readError)
		t.Fatalf(readError.Error())
	}

	// Compare the retrieved user with the original user
	if retrievedUser.Name != user.Name ||
		retrievedUser.Email != user.Email ||
		retrievedUser.Status != user.Status {
		compareError := utils.NewError(
			ErrorLabel, errors.New("retrieved user not match original user"))
		t.Fatalf(compareError.Error())
	}

	// Update the user's status
	var newStatus uint = 2
	statusError := db.Model(&retrievedUser).Update("Status", newStatus).Error
	if statusError != nil {
		statusError = utils.NewError(
			ErrorLabel, statusError)
		t.Fatalf("invalid status: %s", statusError.Error())
	}

	// Read the user again to check the update
	var updatedUser models.User
	readUpdatedUserError := db.First(&updatedUser, user.ID).Error
	if readUpdatedUserError != nil {
		readUpdatedUserError = utils.NewError(ErrorLabel, readUpdatedUserError)
		t.Fatalf("failed to read updated user: %s", readUpdatedUserError.Error())

	}

	// Compare the updated user's status
	if updatedUser.Status != newStatus {
		t.Fatalf("updated user's status does not match the expected value")
	}

	// Delete the user
	deleteError := db.Delete(&updatedUser).Error
	if deleteError != nil {
		deleteError = utils.NewError(ErrorLabel, deleteError)
		t.Fatalf("failed to delete user %s", deleteError.Error())
	}

	// Try to read the user after deletion (should not exist)
	readAfterDeletionError := db.First(&retrievedUser, user.ID).Error
	if readAfterDeletionError == nil {
		readAfterDeletionError = utils.NewError(ErrorLabel, readAfterDeletionError)
		t.Fatalf("expected user to be deleted, but it still exists")
	}
}

func TestCreateDBConnection(t *testing.T) {
	tests := []struct {
		name     string
		dbConfig DBConfig
		envError *utils.CustomError
		want     *utils.CustomError
	}{
		{
			name:     "create connection",
			envError: LoadEnv(".env"),
			dbConfig: DBConfig{
				Host:     os.Getenv("POSTGRES_HOST"),
				Port:     os.Getenv("POSTGRES_PORT"),
				DBName:   os.Getenv("POSTGRES_DB"),
				User:     os.Getenv("POSTGRES_USER"),
				Password: os.Getenv("POSTGRES_PASSWORD"),
				SLLMode:  os.Getenv("POSTGRES_SLL_MODE"),
			},
			want: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !reflect.DeepEqual(tt.envError, tt.want) {
				t.Errorf("CreateDBConnection() envErrOR = %v, want %v", 
          tt.envError, 
          tt.want,
        )
			}

			_, createError := CreateDBConnection(tt.dbConfig)
			if !reflect.DeepEqual(createError, tt.want) {
				t.Errorf("CreateDBConnection() createError = %v, want %v", 
          createError,
          tt.want,
        )
			}
		})
	}
}

func TestLoadEnv(t *testing.T) {
	tests := []struct {
		name string
		want *utils.CustomError
	}{
		{
			name: "env load",
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LoadEnv(".env"); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LoadEnv() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSetupDB(t *testing.T) {
	tests := []struct {
		name string
		want *utils.CustomError
	}{
		{
			name: "setup db",
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, got1 := SetupDB()
			if !reflect.DeepEqual(got1, tt.want) {
				t.Errorf("SetupDB() got1 = %v, want %v", got1, tt.want)
			}
		})
	}
}
