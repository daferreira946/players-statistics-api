package actions

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"players/config"
	"players/models"
)

func SavePlayer(context *gin.Context) {
	var player models.Player

	if err := context.ShouldBindJSON(&player); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	config.DB.Create(&player)

	context.JSON(http.StatusOK, gin.H{
		"message": "Player saved",
		"player":  player.TransposeStructs(),
	})
}
