package main

import (
	"github.com/gin-gonic/gin"
	"student_backend/handlers"
	jwt "student_backend/middleware"
)

func main() {
	router := gin.Default()
	StudentAPI := router.Group("/api/v1")
	{
		userApi := StudentAPI.Group("/user")
		{
			//用户注册
			userApi.POST("/register", handlers.UserRegister)
			//用户登录
			userApi.POST("/login", handlers.UserLogin)
			//根据ID查询用户信息 根据传进来的JWT获取用户信息 需要解析
		}
		bookApi := StudentAPI.Group("/book").Use(jwt.AuthInterceptor())
		{
			//获取全部图书
			bookApi.GET("/", handlers.BookList)
			//根据图书名字获取具体图书信息
			bookApi.POST("/", handlers.BookDetail)
			//删除图书
			bookApi.POST("/del", handlers.BookDeleted)
			//新增图书
			bookApi.POST("/add", handlers.BookAdd)
		}
	}
	router.Run(":3000")
}
