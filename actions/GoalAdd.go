package actions

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"players/config"
	"players/models"
)

func AddGoalToPlayer(context *gin.Context) {
	var player models.Player
	id := context.Params.ByName("id")

	var goals models.ScoreToAdd

	if err := context.ShouldBindJSON(&goals); err != nil {
		context.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": "Validation error",
			"errors":  err.Error(),
		})

		return
	}

	if !goals.ValidateDate() {
		context.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": "Validation error",
			"errors":  "Invalid format date",
		})

		return
	}

	config.DB.First(&player, id)

	if player.ID == 0 {
		context.JSON(http.StatusNotFound, gin.H{})
		return
	}

	var presence models.Presence
	config.DB.Where("player_id = ?", player.ID).Where("date = ?", goals.Date).First(&presence)

	if presence.ID == 0 {
		err := config.DB.Model(&player).Association("Presences").Append(&models.Presence{Date: goals.Date})
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"message": "Could not save the presence",
			})
			return
		}
	}

	for i := 0; i < goals.Value; i++ {
		err := config.DB.Model(&player).Association("Scores").Append(&models.Score{Goal: true, Date: goals.Date})

		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"message": "Could not add goals to player",
			})
			log.Print(err.Error())
			return
		}
	}

	player.Goals += goals.Value
	config.DB.Model(&player).Updates(player)

	context.JSON(http.StatusOK, gin.H{"player": player.TransposeStructs()})
}
