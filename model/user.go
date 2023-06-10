package model

import (
	"gorm.io/gorm"
)

// User is a database model for players
type User struct {
	gorm.Model
	Username  string `json:"username" gorm:"varchar(20);not null"`
	Telephone string `json:"telephone" gorm:"varchar(20);not null;unique"`
	Password  string `json:"password" gorm:"size:255;not null"`
}
