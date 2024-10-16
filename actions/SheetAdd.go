package actions

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"players/config"
	"players/helpers"
	"players/models"
)

func ProcessSheet(context *gin.Context) {
	var sheet models.Sheet

	if err := context.ShouldBindJSON(&sheet); err != nil {
		context.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": "Validation error",
			"errors":  err.Error(),
		})

		return
	}

	date := sheet.Date
	if !helpers.ValidateDate(date) {
		context.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": "Validation error",
			"errors":  "Invalid format date",
		})

		return
	}

	tx := config.DB.Begin()
	for _, stat := range sheet.Stats {
		stat.Date = date
		msg, err := processStats(stat, tx)
		if err != nil {
			tx.Rollback()
			context.JSON(http.StatusBadRequest, gin.H{
				"message": msg,
			})
			return
		}

	}
	tx.Commit()

	context.JSON(http.StatusOK, gin.H{})
}

func processStats(stat models.Stat, tx *gorm.DB) (string, error) {
	var player models.Player
	config.DB.Where("name = ?", stat.Name).First(&player)

	for i := 0; i < stat.Goals; i++ {
		err := tx.Model(&player).Association("Scores").Append(&models.Score{Goal: true, Date: stat.Date})

		if err != nil {
			return "Could not add goals to player", err
		}
	}

	for i := 0; i < stat.Assists; i++ {
		err := tx.Model(&player).Association("Scores").Append(&models.Score{Assist: true, Date: stat.Date})

		if err != nil {
			return "Could not add assists to player", err
		}
	}

	player.Goals += stat.Goals
	player.Assists += stat.Assists

	tx.Model(&player).Updates(player)

	err := tx.Model(&player).Association("Presence").Append(&models.Presence{Date: stat.Date})

	if err != nil {
		return "Could not add presence to player", err
	}

	return "", nil
}
