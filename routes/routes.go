package routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"os"
	"players/actions"
)

func HandleRequests() {
	router := gin.Default()

	configCors(router)

	router.GET("/", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "Health",
		})
	})
	router.GET("/players", actions.GetAllPlayers)
	router.POST("/player", actions.SavePlayer)
	router.GET("/player/:id", actions.GetPlayerById)
	router.PATCH("/player/:id", actions.UpdatePlayer)
	router.DELETE("/player/:id", actions.DeletePlayer)

	router.POST("/player/:id/addGoal", actions.AddGoalToPlayer)
	router.POST("/player/:id/addAssist", actions.AddAssistToPlayer)

	router.GET("/assists", actions.GetAllAssists)
	router.GET("/goals", actions.GetAllGoals)

	router.GET("/top", actions.GetTops)

	router.Run(":8080")
}

func configCors(router *gin.Engine) {
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{os.Getenv("FRONTEND_URL")},
		AllowHeaders: []string{"*"},
		AllowMethods: []string{"*"},
	}))
}
