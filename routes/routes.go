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
		user.GET("/login", userCtr.Login)
		user.GET("/register", userCtr.Register)
		user.GET("/logout", userCtr.Logout)
	}
	ping := router.Group("/ping")
	{
		ping.GET("", pingCtr.Ping)
	}
}



