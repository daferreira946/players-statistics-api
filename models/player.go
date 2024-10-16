package models

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Player struct {
	gorm.Model
	Name      string     `json:"name" binding:"required" gorm:"unique;not null"`
	IsMonthly *bool      `json:"is_monthly" gorm:"default:true"`
	Goals     int        `json:"goals" gorm:"not null"`
	Assists   int        `json:"assists" gorm:"not null"`
	Scores    []Score    `json:"scores" gorm:"not null;foreignkey:PlayerID;constraint:OnDelete:CASCADE"`
	Presence  []Presence `json:"presence" gorm:"not null;foreignkey:PlayerID;constraint:OnDelete:CASCADE"`
}

type PlayersReturnInfo struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	IsMonthly *bool  `json:"is_monthly"`
	Goals     int    `json:"goals"`
	Assists   int    `json:"assists"`
}

func (player *Player) TransposeStructs() PlayersReturnInfo {
	var playerToReturn PlayersReturnInfo
	playerToReturn.ID = player.ID
	playerToReturn.Name = player.Name
	playerToReturn.Goals = player.Goals
	playerToReturn.Assists = player.Assists
	playerToReturn.IsMonthly = player.IsMonthly

	return playerToReturn
}

func (player *Player) AfterDelete(tx *gorm.DB) (err error) {
	tx.Clauses(clause.Returning{}).Where("player_id = ?", player.ID).Delete(&Score{})
	tx.Clauses(clause.Returning{}).Where("player_id = ?", player.ID).Delete(&Presence{})
	return
}
