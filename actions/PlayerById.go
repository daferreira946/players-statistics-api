package actions

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"players/config"
	"players/models"
)

func GetPlayerById(context *gin.Context) {
	var player models.Player
	id := context.Params.ByName("id")

	if err := config.DB.First(&player, id).Error; err != nil {
		context.JSON(http.StatusNotFound, gin.H{})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"player": player.TransposeStructs(),
	})
}
