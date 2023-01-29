package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/richingm/jwt-demo/middleware/jwt"
	"github.com/richingm/jwt-demo/routers/api"
)

func main() {
	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	gin.SetMode(gin.DebugMode)

	r.GET("/auth", api.GetAuth)

	apiv1 := r.Group("/api/v1")
	apiv1.Use(jwt.JWT())
	{
		r.GET("/info", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"data": "info",
			})
		})
	}
	r.Run()
}
