package models

import (
	"time"
)

var RelationStrongSide = map[string]uint{
	"both":  0,
	"left":  1,
	"right": 2,
}

const (
	UStatusDisable    = 1
	UStatusActive     = 2
	UStatusProcessing = 3
)

var UserStatus = map[uint]string{
	UStatusDisable:    "disabled",
	UStatusActive:     "active",
	UStatusProcessing: "processing",
}

const (
	TStatusFinished = 1
	TStatusPending  = 2
	TStatusOverdue  = 3
)

var TaskStatus = map[uint]string{
	TStatusFinished: "finished",
	TStatusPending:  "pending",
	TStatusOverdue:  "overdue",
}

// `json:"-"` Hide from JSON (not exposed)
type User struct {
	ID        uint
	Name      string    `json:"name" gorm:"size:128;not null;"`
	Email     string    `json:"email" gorm:"size:128;not null;"`
	Status    uint      `json:"status"`
	Password  string    `json:"-" gorm:"not null"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoCreateTime:false"`
}

type Role struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	Name        string `json:"name" gorm:"size:128;not null;"`
	Description string `json:"description" gorm:"not null"`
}

type UserRole struct {
	ID     uint `json:"id" gorm:"primaryKey"`
	UserID uint `json:"user_id"`
	RoleID uint `json:"role_id"`
}

type Relation struct {
	ID          uint     `json:"id" gorm:"primaryKey"`
	StrongSide  uint     `json:"strong-side" gorm:"not null"`
	LUserRoleId uint     `json:"l-user-role-id"`
	LUserRole   UserRole `json:"l-user-role" gorm:"foreignKey:LUserRoleId"`
	RUserRoleId uint     `json:"r-user-role-id"`
	RUserRole   UserRole `json:"r-user-role" gorm:"foreignKey:RUserRoleId"`
}

type Task struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Description string    `json:"description" gorm:"not null"`
	AssignedTo  uint      `json:"assigned_to" gorm:"not null"`
	CreatedBy   uint      `json:"created_by"  gorm:"not null"`
	Status      uint      `json:"status"`
	Deadline    time.Time `json:"deadline"    gorm:"not null"`
	UpdatedAt   time.Time `json:"updated_at"  gorm:"autoCreateTime:false"`
}

type Message struct {
	ID       uint   `json:"-" gorm:"primaryKey"`
	Sender   uint   `json:"sender" gorm:"not null"`
	Receiver uint   `json:"receiver" gorm:"not null"`
	Content  string `json:"content" gorm:"not null"`
}
