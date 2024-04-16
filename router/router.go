package router

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/render"
)

func Initialize() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.Render(200, render.JSON{Data: any(gin.H{"message": "pong"})})
	})
	r.Run(":8080")
}
