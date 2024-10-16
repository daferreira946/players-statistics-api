package models

import (
	"gorm.io/gorm"
	"players/helpers"
)

type Score struct {
	gorm.Model
	Goal     bool   `json:"goal" gorm:"not null;default:false"`
	Assist   bool   `json:"assist" gorm:"not null;default:false"`
	PlayerID uint   `json:"player_id" gorm:"not null"`
	Date     string `json:"date" gorm:"default:null"`
}

type ScoreReturnInfo struct {
	ID         uint   `json:"id"`
	Date       string `json:"date"`
	PlayerName string `json:"player_name"`
}

type TopReturnInfo struct {
	Name      string  `json:"name"`
	Quantity  int     `json:"quantity"`
	PerGame   float64 `json:"per_game"`
	IsMonthly bool    `json:"is_monthly"`
}

type ScoreToAdd struct {
	Value     int    `json:"quantity"`
	Date      string `json:"date"`
	IsMonthly bool   `json:"is_monthly"`
}

func (score *ScoreToAdd) ValidateDate() bool {
	return helpers.ValidateDate(score.Date)
}
