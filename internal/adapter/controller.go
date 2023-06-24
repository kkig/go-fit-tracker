package adapter

import (
	"go-fit-tracker/internal/adapter/postgres/models"
	"go-fit-tracker/internal/domain"
	"go-fit-tracker/internal/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// var (
// 	activityRepo 		= repo.Activity{}
// )

func Router() *gin.Engine {
	router := gin.Default()

	router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong!",
		})
	})
	router.POST("/activity", createActivity)

	return router
}

func createActivity(ctx *gin.Context) {
	var input models.ActInput 

	err := ctx.ShouldBindJSON(&input)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	duration, err := strconv.Atoi(input.Duration)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newActivity := domain.ActInput{
		UserID:		input.UserID,
		ExcID:		input.ExcID,
		Duration:	duration,	
		Date:		input.Date,		
	}

	activity := usecase.CreateActivity(newActivity)
	

	ctx.JSON(200, activity)
}