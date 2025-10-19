package router

import (
	"llmcloud/internal/controller"
	"llmcloud/internal/middleware"

	"github.com/gin-gonic/gin"
)

func SetUpRouters(r *gin.Engine, uc *controller.UserController, fc *controller.FileController) {
	// 用户相关路由
	api := r.Group("/api/v1")
	{
		publicUser := api.Group("/users")
		{
			publicUser.POST("/register", uc.Register)
			publicUser.POST("/login", uc.Login)
		}

		auth := api.Group("files")
		auth.Use(middleware.JWTAuth())
		{
			auth.POST("/upload", fc.Upload)
			auth.GET("/download", fc.Download)
			auth.GET("/page", fc.PageList)
			auth.DELETE("/delete", fc.Delete)
			auth.POST("folder", fc.CreateFolder)
			auth.POST("/move", fc.BatchMove)
		}
	}
}
