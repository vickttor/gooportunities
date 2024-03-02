package router

import "github.com/gin-gonic/gin"

// Capital Letter functions is automatically exported.
func Initialize() {
	// Initilize the router with default settings from GIN
	router := gin.Default()

	// Creating a GET Route
	router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.Run(":8080")
}