package routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"players/actions"
	"players/middlewares"
)

func HandleRequests() {
	router := gin.Default()

	configCors(router)

	router.POST("/user/login", actions.Login)
	router.POST("/user/register", middlewares.CheckAuth, actions.CreateUser)

	router.GET("/", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "Health",
		})
	})
	router.GET("/players", actions.GetAllPlayers)
	router.POST("/player", middlewares.CheckAuth, actions.SavePlayer)
	router.GET("/player/:id", actions.GetPlayerById)
	router.PATCH("/player/:id", middlewares.CheckAuth, actions.UpdatePlayer)
	router.DELETE("/player/:id", middlewares.CheckAuth, actions.DeletePlayer)

	router.POST("/player/:id/addGoal", middlewares.CheckAuth, actions.AddGoalToPlayer)
	router.POST("/player/:id/addAssist", middlewares.CheckAuth, actions.AddAssistToPlayer)

	router.GET("/assists", actions.GetAllAssists)
	router.GET("/goals", actions.GetAllGoals)

	router.GET("/top", actions.GetTops)

	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT not set")
	}

	err := router.Run(":" + port)
	if err != nil {
		log.Println("error to start", err.Error())
	}
}

func configCors(router *gin.Engine) {
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{os.Getenv("FRONTEND_URL")},
		AllowHeaders: []string{"*"},
		AllowMethods: []string{"*"},
	}))
}
