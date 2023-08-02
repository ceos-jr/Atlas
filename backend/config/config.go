package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"orb-api/utils"
	"os"
)

type Repository struct {
	DB *gorm.DB
}

type DBConfig struct {
	Host     string
	User     string
	Password string
	DBName   string
	Port     string
	SLLMode  string
}

func LoadEnv(path string) *utils.CustomError {
	loadEnvErrorLabel := "Load env error"
	envError := godotenv.Load(path)

	if envError != nil {
		return utils.NewError(loadEnvErrorLabel, envError)
	}

	return nil
}

func CreateDBConnection(config DBConfig) (*gorm.DB, *utils.CustomError) {
	createDBConnectionErrorLabel := "Connection Error"

	DataSourceName := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		config.Host, config.User, config.Password, config.DBName, config.Port, config.SLLMode,
	)

	connection, dbOpenError := gorm.Open(postgres.Open(DataSourceName), &gorm.Config{})

	if dbOpenError != nil {
		return nil, utils.NewError(createDBConnectionErrorLabel, dbOpenError)
	}

	return connection, nil
}

func SetupDB() (*Repository, *utils.CustomError) {
	setupDBErrorLabel := "SetupDB Error"

	if dbEnvError := LoadEnv(".env"); dbEnvError != nil {
		dbEnvError.AddLabel(setupDBErrorLabel)
		return nil, dbEnvError
	}

	config := DBConfig{
		Host:     os.Getenv("POSTGRES_HOST"),
		Port:     os.Getenv("POSTGRES_PORT"),
		DBName:   os.Getenv("POSTGRES_DB"),
		User:     os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		SLLMode:  os.Getenv("POSTGRES_SLL_MODE"),
	}

	connection, dbConfigError := CreateDBConnection(config)

	if dbConfigError != nil {
		dbConfigError.AddLabel(setupDBErrorLabel)
		return nil, dbConfigError
	}

	return &Repository{
		DB: connection,
	}, nil
}
