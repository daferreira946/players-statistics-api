package actions

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"players/config"
	"players/models"
)

func GetAllPlayers(context *gin.Context) {
	var players []models.PlayersReturnInfo
	name := context.Query("name")

	query := config.DB.Model(&models.Player{})

	if name != "" {
		name = "%" + name + "%"
		query.Where("LOWER(name) like LOWER(?)", name)
	}

	err := query.Order("name asc").Find(&players).Error

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Couldn't get players",
		})

		return
	}

	context.JSON(http.StatusOK, gin.H{
		"players": players,
	})
}
