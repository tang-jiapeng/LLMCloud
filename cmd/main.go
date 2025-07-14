package main

import (
	"llmcloud/config"
	"llmcloud/internal/controller"
	"llmcloud/internal/dao"
	"llmcloud/internal/database"
	"llmcloud/internal/router"
	"llmcloud/internal/service"
	"time"

	"github.com/gin-contrib/cors"
)

func main() {
	config.InitConfig()

	db, _ := database.InitDB()

	userDao := dao.NewUserDao(db)
	userService := service.NewUserService(userDao)
	userController := controller.NewUserController(userService)

	r := router.SetUserRouter(userController)

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},                                                 // 允许所有域名
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},  // 允许的HTTP方法
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"}, // 允许的请求头
		ExposeHeaders:    []string{"Content-Length"},                                    // 暴露的响应头
		AllowCredentials: true,                                                          // 允许携带凭证（如Cookie）
		MaxAge:           12 * time.Hour,
	}))

	r.Run(":8080")
}
