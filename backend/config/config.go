package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"orb-api/models"
	"orb-api/repositories"
	"os"
)

type DBConfig struct {
	Host     string
	User     string
	Password string
	DBName   string
	Port     string
	SLLMode  string
}

func LoadEnv(path string) error {
	return godotenv.Load(path)
}

func CreateDBConnection(config DBConfig) (*gorm.DB, error) {
	DataSourceName := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		config.Host,
		config.User,
		config.Password,
		config.DBName,
		config.Port,
		config.SLLMode,
	)

	connection, dbOpenError := gorm.Open(
		postgres.Open(DataSourceName), &gorm.Config{},
	)

	if dbOpenError != nil {
		return nil, dbOpenError
	}

	return connection, nil
}

func MigrateDB(database *gorm.DB) error {
	// Skip migration
	if os.Getenv("MIGRATE") == "false" {
		return nil
	}

	return database.Migrator().AutoMigrate(
		&models.User{},
		&models.Role{},
		&models.UserRole{},
		&models.Relation{},
		&models.Task{},
		&models.Message{},
	)
}

func SetupDB() (*repository.Repository, error) {
	if dbEnvError := LoadEnv(".env"); dbEnvError != nil {
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

	connection, connectionError := CreateDBConnection(config)

	if connectionError != nil {
		return nil, connectionError
	}

	if migrationError := MigrateDB(connection); migrationError != nil {
		return nil, migrationError
	}

	return repository.SetupRepository(connection), nil
}

func CloseDB(repository *repository.Repository) error {
	sqlDB, _ := repository.DB.DB()

	return sqlDB.Close()
}

func PingDB(repository *repository.Repository) error {
	sqlDB, _ := repository.DB.DB()

	return sqlDB.Ping()
}
