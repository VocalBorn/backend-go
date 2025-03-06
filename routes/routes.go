package routes

import (
	"vocalborn/backend-go/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine)  {
	router.GET("/ping", handlers.Ping)
}