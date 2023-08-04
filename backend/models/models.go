package models 

import (
	"gorm.io/gorm"
	"time"
)

// `json:"-"` Hide from JSON (not exposed)
type User struct {
	gorm.Model 					`json:"gorm_._model"`
	Name			string 		`json:"name" gorm:"size:128;not null;"`
	Email			string 		`json:"email" gorm:"size:128;not null;"`
	Status    uint		  `json:"status"`
  Password  string	  `json:"-" gorm:"not null"` 
	UpdatedAt time.Time `json:"updated_at" gorm:"autoCreateTime:false"`
}

type Role struct {
	ID				  uint		`json:"id" gorm:"primaryKey"`
	Name			  string	`json:"name" gorm:"size:128;not null;"`
	Description	string	`json:"description" gorm:"not null"`
}

type UserRole struct {
	ID		  uint	`json:"id" gorm:"primaryKey"`
	UserID 	uint	`json:"user_id"`
	RoleID 	uint 	`json:"role_id"`
}

// side -> embedded in Relation
type side struct {
	ID uint
	PositionType string
}

type Relation struct {
	ID 			uint	`json:"id" gorm:"primaryKey"`
  Right		side	`json:"right" gorm:"embedded; embeddedPrefix:right_"`
	Left		side	`json:"left" gorm:"embedded; embeddedPrefix:left_"`
}

type Task struct {
	gorm.Model
  Description 	string	  `json:"description" gorm:"not null"`
  AssignedTo		uint		  `json:"assigned_to" gorm:"not null"`
  CreatedBy		  uint		  `json:"created_by"  gorm:"not null"`
	Status			  string	  `json:"status"`
	UpdatedAt 		time.Time `json:"updated_at" gorm:"autoCreateTime:false"`
}

type Message struct {		// CreatedAt <-> SendAt
  gorm.Model	
  Sender 		uint    `json:"sender" gorm:"not null"`
  Receiver	uint 		`json:"receiver" gorm:"not null"`
  Content		string 	`json:"content" gorm:"not null"`
}
