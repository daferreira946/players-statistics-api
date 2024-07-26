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

	log.Println(goals)

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

	player.Goals += goals.Value

	config.DB.Model(&player).Updates(player)

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

	context.JSON(http.StatusOK, gin.H{"player": player.TransposeStructs()})
}
