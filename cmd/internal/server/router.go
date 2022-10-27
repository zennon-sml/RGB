package router

import (
	"github.com/gin-gonic/gin"
)

func setRouter() *gin.Engine {
	router := gin.Default()

	api := router.Group("/api")
	{
		api.GET("/gello", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{"msg": "gello gin"})
		})
	}
	router.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(400, gin.H{"msg": "error handling the routes"})
	})
	return router
}
