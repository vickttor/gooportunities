package router

import "github.com/gin-gonic/gin"

// Capital Letter functions is automatically exported.
func Initialize() {
	// Initilize the router with default settings from GIN
	router := gin.Default()

	// Initializing the routes
	initializeRoutes(router)

	// Running the server
	router.Run(":8080")
}