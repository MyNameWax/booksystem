package config

import (
	"fmt"
	jwtUtil "student_backend/util"

	"github.com/gin-gonic/gin"
)

func AuthInterceptor() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//获取请求头
		token := ctx.GetHeader("Authorization")
		fmt.Println("token信息", token)
		if token == "" {
			ctx.JSON(401, gin.H{
				"code":    401,
				"message": "未登录",
			})
			ctx.Set("error", "error")
			return
		}
		result := jwtUtil.ParseJwt(token)
		if result == "" {
			ctx.JSON(401, gin.H{
				"code":    401,
				"message": "身份验证失败,请重新登录",
			})
			ctx.Set("error", "error")
			return
		}
		ctx.Next()
	}
}
