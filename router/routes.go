package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {
	// Pointer:

	// It's a wait to reference to a value in memory. Instead of pass the value
	// through the memory slots, we pass the address of the first variable 
	// which stores the value, this way we save memory and increase performance.

	// Grouping endpoints
	v1 := router.Group("/api/v1")

	{	
		v1.GET("/opening", func (ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"msg": "Get Opening",
			})
		})
		
		v1.POST("/opening", func (ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"msg": "POST Opening",
			})
		})	
		
		v1.DELETE("/opening", func (ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"msg": "DELETE Opening",
			})
		})	
		
		v1.PUT("/opening", func (ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"msg": "PUT Opening",
			})
		})	

		v1.GET("/openings", func (ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"msg": "Get Openings",
			})
		})	
	}
}