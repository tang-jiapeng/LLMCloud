package middleware

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

const (
	AuthHeaderKey  = "Authorization"
	AuthBearerType = "Bearer"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取 Authorization 头
		authHeader := c.GetHeader(AuthHeaderKey)
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"code":    http.StatusUnauthorized,
				"message": "需要认证令牌",
			})
			return
		}

		// 分割 Bearer 和令牌
		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || parts[0] != AuthBearerType {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"code":    http.StatusUnauthorized,
				"message": "令牌格式错误，应为 Bearer <token>",
			})
			return
		}

		// 解析验证令牌
		tokenString := parts[1]
		claims, err := ParseToken(tokenString)

		if err != nil {
			status := http.StatusUnauthorized
			message := "无效令牌"

			switch {
			case errors.Is(err, jwt.ErrTokenExpired):
				message = "令牌已过期"
				status = http.StatusForbidden
			case errors.Is(err, jwt.ErrTokenMalformed):
				message = "令牌格式错误"
			case errors.Is(err, jwt.ErrSignatureInvalid):
				message = "签名验证失败"
			}

			c.AbortWithStatusJSON(status, gin.H{
				"code":    status,
				"message": message,
			})
			return
		}

		// 将用户ID存入上下文
		c.Set("user_id", claims.UserID)
		c.Next()
	}
}
