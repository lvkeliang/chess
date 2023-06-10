package model

import "gorm.io/gorm"

// Board is a database model for chess boards
type Board struct {
	gorm.Model
	Name       string `gorm:"not null"`
	WhiteID    uint   `gorm:"not null"` // the user ID of the white player
	BlackID    uint   `gorm:"not null"` // the user ID of the black player
	WhiteReady bool   `gorm:"not null"`
	BlackReady bool   `gorm:"not null"`
	Playing    bool   `gorm:"not null"`
	AudienceID string `gorm:"default:'{}'"`
	State      []byte `gorm:"type:text(384);not null"` // the state of the board, represented by 64 characters
	Turn       bool   `gorm:"not null"`                // the turn of the players, true for white, false for black
}
