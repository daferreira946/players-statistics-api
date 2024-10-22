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

	query.Joins("join players on scores.player_id = players.id").Select("count(scores.*) as quantity, scores.date as date, players.name as player_name")

	err := query.Group("players.name, scores.date").Order("date desc").Order("player_name asc").Find(&goals).Error

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
