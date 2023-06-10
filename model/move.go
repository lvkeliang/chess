package model

import "gorm.io/gorm"

// Move is a database model for chess moves
type Move struct {
	gorm.Model
	BoardID uint   `gorm:"not null"`              // the board ID of the move
	UserID  uint   `gorm:"not null"`              // the user ID of the move maker
	From    string `gorm:"type:char(2);not null"` // the from position of the move, such as "e2"
	To      string `gorm:"type:char(2);not null"` // the to position of the move, such as "e4"
}
