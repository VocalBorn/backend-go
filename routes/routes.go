package routes

import (
	"vocalborn/backend-go/handlers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine)  {
	user := router.Group("/user")
	{
		user.GET("/login", handlers.Login)
		user.GET("/register", handlers.Register)
		user.GET("/logout", handlers.Logout)
	}
}

