package router

import (
	"github.com/gin-gonic/gin"
	"github.com/mcondero/goportunidades/handler"
)

func initializeRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	{
		v1.GET("/opening", handler.GetOpeningHandler)
		v1.POST("/opening", handler.CreateOpeningHandler)
		v1.DELETE("/opening", handler.DeleteOpeningHandler)
		v1.PUT("/opening", handler.PutOpeningHandler)
		v1.GET("/openings", handler.ListOpeningsHandler)
	}
}
