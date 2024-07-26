package actions

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
	"net/http"
	"players/config"
	"players/models"
)

func DeletePlayer(context *gin.Context) {
	var player models.Player
	id := context.Params.ByName("id")

	if err := config.DB.First(&player, id).Error; err != nil {
		context.JSON(http.StatusNotFound, gin.H{})
		return
	}

	if err := config.DB.Select(clause.Associations).Delete(&player).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not delete player associations.",
		})
		return
	}

	config.DB.Delete(&player, id)
	context.JSON(http.StatusOK, gin.H{"message": "Player deleted"})
}
