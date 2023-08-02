package config

import (
	"gorm.io/gorm"
	"time"
)

// Model -> gorm.Model definition
type Model struct{
	Id 				uint		`json:"id" gorm:"primaryKey"` // Primary key
	CreatedAt		time.Time 	`json:"created_at" gorm:"autoCreateTime:true"`
}

type User struct {
	gorm.Model 					`json:"gorm_._model"`
	Name			string 		`json:"name" gorm:"size:128;not null;"`
	Email			string 		`json:"email" gorm:"size:128;not null;"`
	// Avatar    string    `json:"avatar"`  ( Define later )
	Status    		uint		`json:"status"`
	Password  		string		`json:"-"` // Hide from JSON (not exposed)
	UpdatedAt 		time.Time 	`json:"updated_at" gorm:"autoCreateTime:false"`
}

type Role struct {
	Id				uint		`json:"id" gorm:"primaryKey"`
	Name			string		`json:"name" gorm:"size:128;not null;"`
	Description		string		`json:"description" gorm:""`
}

type UserRole struct {
	Id				uint		`json:"id" gorm:"primaryKey"`
	UserId 			uint		`json:"user_id"`
	RoleId 			uint 		`json:"role_id"`
}

// side -> embedded in Relation
type side struct {
	id uint
	positionType string
}

type Relation struct {
	Id 				uint		`json:"id" gorm:"primaryKey"`
	SideA			side		`json:"side-a" gorm:"embedded"`
	SideB			side		`json:"side-b" gorm:"embedded"`
}

type Task struct {
	gorm.Model
	Description 	string		`json:"description"`
	AssignedTo		uint		`json:"assigned_to"`
	CreatedBy		uint		`json:"created_by"`
	Status			string		`json:"status"`
	UpdatedAt 		time.Time 	`json:"updated_at" gorm:"autoCreateTime:false"`
}

type Message struct {		// CreatedAt <-> SendAt
	Id 				uint		`json:"id" gorm:"primaryKey"`
	Sender 			UserRole 	`json:"sender"`
	Receiver		uint 		`json:"receiver"`
	Data			string 		`json:"data"`
}
