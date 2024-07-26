package actions

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"players/config"
	"players/models"
)

func UpdatePlayer(context *gin.Context) {
	var player models.Player
	id := context.Params.ByName("id")
	config.DB.First(&player, id)

	if err := context.ShouldBindJSON(&player); err != nil {
		context.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": "Validation error",
			"errors":  err.Error(),
		})

		return
	}

	if err := config.DB.Model(&player).Updates(player).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Couldn't update the player",
			"error":   err.Error(),
		})

		return
	}

	context.JSON(http.StatusOK, gin.H{"player": player.TransposeStructs()})
}
