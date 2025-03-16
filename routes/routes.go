package routes

import (
	"vocalborn/backend-go/controllers"
	"github.com/gin-gonic/gin"
	)

func SetupRoutes(router *gin.Engine)  {
	pingCtr := controllers.NewPingController()
	userCtr := controllers.NewUserController()

	user := router.Group("/user")
	{
		user.POST("/login", userCtr.Login)
		user.POST("/register", userCtr.Register)
		user.POST("/logout", userCtr.Logout)
	}
	ping := router.Group("/ping")
	{
		ping.GET("", pingCtr.Ping)
	}
}



