package main

import (
	_ "vocalborn/backend-go/docs"
	"vocalborn/backend-go/models"
	"vocalborn/backend-go/routes"
	"vocalborn/backend-go/utils"

	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

//	@title		Vocalborn API
//	@version	1.0
func main()  {
	// 連接資料庫
	db := utils.ConnectDB()
	db.DryRun = false // 設定為true時，GORM會在執行時顯示SQL指令，但不會真的執行
	db.AutoMigrate(&models.Account{}, &models.User{}) // 自動建立資料表

	// 建立Gin伺服器
	router := gin.Default()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.SetTrustedProxies([]string{"localhost"})
	routes.SetupRoutes(router)
	router.Run(":8080")
}