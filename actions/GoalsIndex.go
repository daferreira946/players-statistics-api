package actions

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"players/config"
	"players/models"
)

func GetAllGoals(context *gin.Context) {
	var goals []models.ScoreReturnInfo
	monthYear := context.Query("month_year")

	query := config.DB.Model(&models.Score{}).Where("scores.goal = ?", true)

	if monthYear != "" {
		query.Where("left(scores.date, 7) = ?", monthYear)
	}

	query.Joins("join players on scores.player_id = players.id").Select("scores.id as id, scores.date as date, players.name as player_name")

	err := query.Find(&goals).Error

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Couldn't get goals",
			"erro":    err.Error(),
		})

		return
	}

	context.JSON(http.StatusOK, gin.H{
		"goals": goals,
	})
}
