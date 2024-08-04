package router

import (
	"github.com/gin-gonic/gin"
	"github.com/tdayko/simple-api-go/handler"
)

func initializeRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1/")
	{
		v1.GET("/openings", handler.ShowOpeningsHandler)
		v1.POST("/openings", handler.CreateOpeningHandler)
		v1.DELETE("/openings", handler.DeleteOpeningByIdOpeningHandler)
		v1.PUT("/openings", handler.UpdateOpeningHandler)
		v1.GET("/openings", handler.ListOpeningsHandler)
	}
}
