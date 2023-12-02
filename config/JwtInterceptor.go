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
				"message": "未登录,请先登录",
			})
			ctx.Abort()
			return
		}
		result, _ := jwtUtil.ParseJwt(token)
		if result == nil {
			ctx.JSON(401, gin.H{
				"code":    401,
				"message": "未登录,请先登录",
			})
			ctx.Abort()
			return
		}
		fmt.Println("登录的用户:", result.Name)
		ctx.Next()
	}
}
