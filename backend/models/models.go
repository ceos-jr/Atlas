package models 

import (
	"gorm.io/gorm"
	"time"
)

// Model -> gorm.Model definition
type Model struct{
	Id 				  uint		  `json:"id" gorm:"primaryKey"` 
	CreatedAt		time.Time `json:"created_at" gorm:"autoCreateTime:true"`
}

// `json:"-"` Hide from JSON (not exposed)
type User struct {
	gorm.Model 					`json:"gorm_._model"`
	Name			string 		`json:"name" gorm:"size:128;not null;"`
	Email			string 		`json:"email" gorm:"size:128;not null;"`
	Status    uint		  `json:"status"`
	Password  string	  `json:"-"` 
	UpdatedAt time.Time `json:"updated_at" gorm:"autoCreateTime:false"`
}

type Role struct {
	Id				  uint		`json:"id" gorm:"primaryKey"`
	Name			  string	`json:"name" gorm:"size:128;not null;"`
	Description	string	`json:"description" gorm:""`
}

type UserRole struct {
	Id		  uint	`json:"id" gorm:"primaryKey"`
	UserId 	uint	`json:"user_id"`
	RoleId 	uint 	`json:"role_id"`
}

// side -> embedded in Relation
type side struct {
	id uint
	positionType string
}

type Relation struct {
	Id 			uint	`json:"id" gorm:"primaryKey"`
	SideA		side	`json:"side-a" gorm:"embedded"`
	SideB		side	`json:"side-b" gorm:"embedded"`
}

type Task struct {
	gorm.Model
	Description 	string	  `json:"description"`
	AssignedTo		uint		  `json:"assigned_to"`
	CreatedBy		  uint		  `json:"created_by"`
	Status			  string	  `json:"status"`
	UpdatedAt 		time.Time `json:"updated_at" gorm:"autoCreateTime:false"`
}

type Message struct {		// CreatedAt <-> SendAt
  gorm.Model	
	Sender 		uint    `json:"sender"`
	Receiver	uint 		`json:"receiver"`
	Data			string 	`json:"data"`
}
