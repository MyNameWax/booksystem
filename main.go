package main

import (
	app "student_backend/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	StudentAPI := router.Group("/api/v1")
	{
		userApi := StudentAPI.Group("/user")
		{
			//用户注册
			userApi.POST("/register", app.UserRegister)
			//用户登录
			userApi.POST("/login", app.UserLogin)
			//根据ID查询用户信息 根据传进来的JWT获取用户信息 需要解析
		}
		bookApi := StudentAPI.Group("/book")
		{
			//获取全部图书
			bookApi.GET("/", app.BookList)
			//根据图书名字获取具体图书信息
			bookApi.POST("/", app.BookDetail)
			//删除图书
			bookApi.POST("/del", app.BookDeleted)
			//修改图书
			//新增图书
		}
	}
	router.Run(":3000")
}
