package config

import (
	"fmt"
	"log"
	"os"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Repository struct {
  DB *gorm.DB
}

type DBConfig struct {
  Host      string 
  User      string 
  Password  string
  DBName    string 
  Port      string 
  SLLMode   string
}

func LoadEnv() error {
  error := godotenv.Load(".env")

  if error != nil {
    return error
  }

  return nil
}

func CreateDBConnection(config DBConfig) (*gorm.DB, error) {
  DataSourceName := fmt.Sprintf(
    "host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", 
    config.Host, config.User, config.Password, config.DBName, config.Port, config.SLLMode, 
  )
  
  connection, error := gorm.Open(postgres.Open(DataSourceName), &gorm.Config{})

  if error != nil {
    return nil, error
  }
 
  return connection, nil
} 

func SetupDB() (*Repository) {
  if error := LoadEnv(); error != nil {
    log.Fatalf("Failed to load .env file: %v\n", error)
  }
  
  config := DBConfig{
    Host:     os.Getenv("POSTGRES_HOST"),
    Port:     os.Getenv("POSTGRES_PORT"),
    DBName:   os.Getenv("POSTGRES_DB"),
    User:     os.Getenv("POSTGRES_USER"),
    Password: os.Getenv("POSTGRES_PASSWORD"),
    SLLMode:  os.Getenv("POSTGRES_SLL_MODE"),
  }

  connection, error := CreateDBConnection(config) 

  if error != nil {
    log.Fatalf("Failed to create a database connection: %v\n", error)
  }

  return &Repository{
    DB: connection,
  }
} 

