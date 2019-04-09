package model

import (
	"time"

	"github.com/jinzhu/gorm"

	"github.com/google/uuid"
)

type Data struct {
	ResourceID string `gorm:"column:resourceId" json:"resource_id"`
}

type Model struct {
	ID        *uuid.UUID `gorm:"primary_key;type:char(36)" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}

type User struct {
	Model
	Username string `json:"username"`
	// Name   string
	SessionID  string  `json:"session_id"`
	Expense    int     `json:"expense"` // Expense Percentage
	Salary     float64 `json:"salary"`
	SalaryDate int     `json:"salary_date"`
}

type Expense struct {
	Model
	UserID uuid.UUID `json:"user_id"`
	User   User      `gorm:"foreignkey:ID;association_foreignkey:UserID" json:"-"`
	Type   string    `sql:"type:ENUM('food', 'transport', 'leisure', 'shopping', 'misc', 'saving', 'morgage', 'charity', 'others')" json:"type"`
	Amount float64   `json:"amount"`
}

func (r *User) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("ID", uuid.New())
	return nil
}

func (p *Expense) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("ID", uuid.New())
	return nil
}
