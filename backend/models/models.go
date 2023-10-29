package models

import (
	"time"
)

//revive:disable:line-length-limit

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

type User struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
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
	LUserRoleID uint     `json:"l-user-role-id"`
	LUserRole   UserRole `json:"l-user-role" gorm:"foreignKey:LUserRoleID"`
	RUserRoleID uint     `json:"r-user-role-id"`
	RUserRole   UserRole `json:"r-user-role" gorm:"foreignKey:RUserRoleID"`
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
	ID        uint      `json:"-" gorm:"primaryKey"`
	Sender    uint      `json:"sender" gorm:"not null"`
	Receiver  uint      `json:"receiver" gorm:"not null"`
	Content   string    `json:"content" gorm:"not null"`
	Timestamp time.Time `json:"timestamp"`
}

type Project struct {
	ID        uint      `json: "id" gorm:"primaryKey"`
	Name      string    `json: "name" gorm:"size:128;not null;"`
	SectorID  uint 		`json: "id" gorm:"foreingKey; not null";`
	AdmID     uint		`json: "id" gorm:"foreingKey; not null";`
}

type TasksProject struct {
	ID        uint     `json:"id" gorm:"primaryKey"`
	TaskID	  uint     `json:"task_id"`
	ProjectID uint     `json:"project_id"` 
}

type UsersProject struct {
	ID        uint     `json:"id" gorm:"primaryKey"`
	UserID	  uint     `json:"user_id"`
	ProjectID uint     `json:"project_id"` 
}


//Criar model de setores
//Conexão com mensagens do whatspp em projects
//Calendário
//Avisos: model (notificação)