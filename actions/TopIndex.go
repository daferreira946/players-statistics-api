package actions

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"players/config"
	"players/models"
	"strconv"
	"strings"
)

func GetTops(context *gin.Context) {
	var goals []models.TopReturnInfo
	var assists []models.TopReturnInfo
	limitString := context.Query("limit")
	monthYear := context.Query("month_year")
	onlyMonthly := context.Query("only_monthly")
	query := config.DB.Model(&models.Score{})

	if limitString != "" {
		limit, err := strconv.Atoi(limitString)

		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"message": "Limit must be a number",
			})

			return
		}

		query.Limit(limit)
	}

	if onlyMonthly != "" {
		onlyMonthlyBool, err := strconv.ParseBool(onlyMonthly)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"message": "Only monthly must be a true or false",
			})

			return
		}

		if onlyMonthlyBool {
			query.Where("players.is_monthly = ?", onlyMonthlyBool)
		}
	}

	subquery := "select count(*) from presences where presences.player_id = players.id"

	if monthYear != "" {
		split := strings.Split(monthYear, "-")

		query.Where("date_part('month', to_date(scores.date, 'YYYY-MM-DD')) = ?", split[1])
		subquery += fmt.Sprintf(" AND date_part('month', to_date(presences.date, 'YYYY-MM-DD')) = '%s'", split[1])
		query.Where("date_part('year', to_date(scores.date, 'YYYY-MM-DD')) = ?", split[0])
		subquery += fmt.Sprintf(" AND date_part('year', to_date(presences.date, 'YYYY-MM-DD')) = '%s'", split[0])
	}

	query.Joins("join players on scores.player_id = players.id").Select(fmt.Sprintf("count(scores.*) as quantity, players.name as name, players.is_monthly as is_monthly, count(scores.*) / NULLIF((%s), 0)::float as per_game", subquery))
	query.Group("players.id, players.name, players.is_monthly").Order("quantity desc").Order("per_game desc")

	goalsQuery := query.Session(&gorm.Session{})

	assistsQuery := query.Session(&gorm.Session{})

	err := goalsQuery.Debug().Where("goal = ?", true).Find(&goals).Error

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Couldn't get top goals",
			"erro":    err.Error(),
		})

		return
	}

	err = assistsQuery.Debug().Where("assist = ?", true).Find(&assists).Error

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Couldn't get top assists",
			"erro":    err.Error(),
		})

		return
	}

	context.JSON(http.StatusOK, gin.H{
		"top_goals":   goals,
		"top_assists": assists,
	})
}
