package model

import "gorm.io/gorm"

type Ready struct {
	gorm.Model
	BoardID uint `gorm:"not null"` // the board ID of the move
	UserID  uint `gorm:"not null"` // the user ID of the move maker
	Ready   bool `gorm:"not null"`
}
