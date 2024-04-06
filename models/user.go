package models

import "time"

type UserType int8

const (
	ADMIN UserType = iota
	CUSTOMER
)

// User is model for implement user account.
type User struct {
	ID        int       `gorm:"primaryKey" json:"id,omitempty"`
	FirstName string    `gorm:"size:50" json:"first_name,omitempty"`
	LastName  string    `gorm:"size:50" json:"last_name,omitempty"`
	Username  string    `gorm:"size:50;uniqueIndex" json:"username,omitempty"`
	Email     string    `gorm:"size:50;unique" json:"email,omitempty"`
	Password  string    `gorm:"size:128" json:"-,omitempty"`
	Type      UserType  `gorm:"default:1" json:"type,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}
