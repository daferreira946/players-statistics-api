package models

import (
	"gorm.io/gorm"
	"log"
	"regexp"
	"time"
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
	Name      string `json:"name"`
	Quantity  int    `json:"quantity"`
	IsMonthly bool   `json:"is_monthly"`
}

type ScoreToAdd struct {
	Value     int    `json:"quantity"`
	Date      string `json:"date"`
	IsMonthly bool   `json:"is_monthly"`
}

func (score *ScoreToAdd) ValidateDate() bool {
	pattern := regexp.MustCompile("^[0-9]{4}-[0-9]{2}-[0-9]{2}$")

	validString := pattern.MatchString(score.Date)

	log.Println(score.Date)

	if !validString {
		return false
	}

	_, err := time.Parse("2006-01-02", score.Date)
	if err != nil {
		return false
	}

	return true
}
