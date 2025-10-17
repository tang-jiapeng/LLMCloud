package router

import (
	"llmcloud/internal/controller"
	"llmcloud/internal/middleware"

	"github.com/gin-gonic/gin"
)

func SetUpRouters(r *gin.Engine, uc *controller.UserController) {
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
			//
		}
	}
}
