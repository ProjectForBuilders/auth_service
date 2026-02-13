package models

import (
	"time"
)

type User struct {
	ID           string  `json:"id"`
	Email        string  `json:"email"`
	Phone        *string `json:"phone"`
	PasswordHash string  `json:"-"`

	FirstName string  `json:"first_name"`
	LastName  string  `json:"last_name"`
	AvatarUrl *string `json:"avatar_url"`

	IsActive   bool       `json:"is_active"`
	IsVerified bool       `json:"is_verified"`
	LastLogin  *time.Time `json:"last_login"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}
