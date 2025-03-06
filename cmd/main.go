package main

import (
	"vocalborn/backend-go/routes"
	"github.com/gin-gonic/gin"
)

func main()  {
	router := gin.Default()
	// 設定信任的代理伺服器
	router.SetTrustedProxies([]string{"localhost"})

	// 設定路由
	routes.SetupRoutes(router)

	// 啟動伺服器
	router.Run(":8080")
}