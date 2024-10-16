package models

import (
	"gorm.io/gorm"
)

type Presence struct {
	gorm.Model
	PlayerID uint   `json:"player_id" gorm:"not null"`
	Date     string `json:"date" gorm:"not null"`
}
