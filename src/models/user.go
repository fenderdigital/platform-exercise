package models

import (
	"time"
)

type User struct {
	ID        uint   `gorm:"primarykey"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"-"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
