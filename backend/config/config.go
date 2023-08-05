package config

import (
	"fmt"
	"orb-api/models"
	"os"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
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

func SetupDB() (*Repository, error) {
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

	return &Repository{
		DB: connection,
	}, nil
}

func (repository *Repository) CloseDB() error {
  sqlDB, _ := repository.DB.DB()
  
  return sqlDB.Close() 
}  

func (repository *Repository) PingDB() error {
  sqlDB, _ := repository.DB.DB()
  
  return sqlDB.Ping() 
}
