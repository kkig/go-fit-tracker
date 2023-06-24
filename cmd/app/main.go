package main

import (
	"go-fit-tracker/internal/adapter"
)

func main() {
	router := adapter.Router()
	// httpPort := os.Getenv("HTTP_PORT")
	// if httpPort == "" {
	// 	httpPort = "8080"
	// }

	// router := gin.Default()

	// v1 := router.Group("/v1")
	// {
	// 	v1.GET("/ping", func(ctx *gin.Context) {
	// 		ctx.JSON(200, gin.H{
	// 			"message": "pong!",
	// 		})
	// 	})
	// }

	router.Run()
}