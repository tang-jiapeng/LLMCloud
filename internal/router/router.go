package router

import (
	"llmcloud/internal/controller"

	"github.com/gin-gonic/gin"
)

func SetUserRouter(uc *controller.UserController) *gin.Engine {
	r := gin.Default()
	// 用户相关路由
	api := r.Group("/api/v1")
	{
		userGroup := api.Group("/users")
		{
			userGroup.POST("/register", uc.Register)
			userGroup.POST("/login", uc.Login)
		}
	}
	return r
}
