package actions

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"players/config"
	"players/models"
)

func AddAssistToPlayer(context *gin.Context) {
	var player models.Player
	id := context.Params.ByName("id")

	var assists models.ScoreToAdd

	if err := context.ShouldBindJSON(&assists); err != nil {
		context.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": "Validation error",
			"errors":  err.Error(),
		})

		return
	}

	if !assists.ValidateDate() {
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
	config.DB.Where("player_id = ?", player.ID).Where("date = ?", assists.Date).First(&presence)

	if presence.ID == 0 {
		err := config.DB.Model(&player).Association("Presences").Append(&models.Presence{Date: assists.Date})
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"message": "Could not save the presence",
			})
			return
		}
	}

	player.Assists += assists.Value
	config.DB.Model(&player).Updates(player)

	for i := 0; i < assists.Value; i++ {
		err := config.DB.Model(&player).Association("Scores").Append(&models.Score{Assist: true, Date: assists.Date})

		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"message": "Could not add assists to player",
			})
			log.Print(err.Error())
			return
		}
	}

	player.Assists += assists.Value
	config.DB.Model(&player).Updates(player)

	context.JSON(http.StatusOK, gin.H{"player": player.TransposeStructs()})
}
