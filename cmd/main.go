package main

import (
	"llmcloud/config"
	"llmcloud/internal/controller"
	"llmcloud/internal/dao"
	"llmcloud/internal/database"
	"llmcloud/internal/middleware"
	"llmcloud/internal/router"
	"llmcloud/internal/service"

	"github.com/gin-gonic/gin"
)

func main() {
	config.InitConfig()

	db, _ := database.InitDB()

	userDao := dao.NewUserDao(db)
	userService := service.NewUserService(userDao)
	userController := controller.NewUserController(userService)

	r := gin.Default()
	// 配置跨域
	r.Use(middleware.SetupCORS())
	router.SetUpRouters(r, userController)

	r.Run(":8080")
}
