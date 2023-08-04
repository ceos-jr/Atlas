package config

import (
	"fmt"
	"orb-api/models"
	"orb-api/utils"
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

func LoadEnv(path string) *utils.CustomError {
	loadEnvErrorLabel := "Load env error"
	
	if envError := godotenv.Load(path); envError != nil {
		return utils.NewError(loadEnvErrorLabel, envError)
	}

	return nil
}

func CreateDBConnection(config DBConfig) (*gorm.DB, *utils.CustomError) {
	connectionErrorLabel := "Connection Error"

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
		return nil, utils.NewError(connectionErrorLabel, dbOpenError)
	}

	return connection, nil
}

func MigrateDB(database *gorm.DB) *utils.CustomError {
  migrateDBErrorLabel := "Migration Error"
  
  // Skip migration
  if os.Getenv("MIGRATE") == "false" {
    return nil
  }

  migrationError := database.Migrator().AutoMigrate(
    &models.User{}, 
    &models.Role{},
    &models.UserRole{},
    &models.Relation{},
    &models.Task{},
    &models.Message{},
  )

  if migrationError != nil {
    return utils.NewError(migrateDBErrorLabel, migrationError) 
  }

  return nil 
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

	connection, connectionError := CreateDBConnection(config)

	if connectionError != nil {
		connectionError.AddLabel(setupDBErrorLabel)
		return nil, connectionError
	}

  migrationError := MigrateDB(connection) 

  if migrationError != nil {
    migrationError.AddLabel(setupDBErrorLabel)
    return nil, migrationError 
  }

	return &Repository{
		DB: connection,
	}, nil
}

func (repository *Repository) CloseDB() (*utils.CustomError) {
  closeDBErrorLabel := "Close DB error" 

  sqlDB, _ := repository.DB.DB()
  
  if closeError := sqlDB.Close(); closeError != nil {
    return utils.NewError(closeDBErrorLabel, closeError) 
  }
  
  return nil
}  

func (repository *Repository) PingDB() (*utils.CustomError) {
  pingDBErrorLabel := "Ping DB error" 

  sqlDB, _ := repository.DB.DB()
  
  if pingError := sqlDB.Ping(); pingError != nil {
    return utils.NewError(pingDBErrorLabel, pingError) 
  }
  
  return nil
}
