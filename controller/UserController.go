package controller

import (
	"net/http"
	dbUtil "student_backend/util"
	jwtUtil "student_backend/util"

	"github.com/gin-gonic/gin"
)

var db = dbUtil.InitDB()

type Users struct {
	Id       int
	Name     string
	Password string
}
type UsersDTO struct {
	Name  string
	Token string
}

func UserLogin(context *gin.Context) {
	var user = Users{}
	if err := context.ShouldBindJSON(&user); err != nil {
		context.JSON(200, gin.H{
			"code": 400,
			"msg":  "请求参数错误",
		})
		return
	}
	result := db.Where("name =?", user.Name).Where("password =?", user.Password).First(&user)
	if result.RowsAffected == 0 {
		context.JSON(http.StatusOK, gin.H{
			"code": 10001,
			"msg":  "账号或密码错误",
		})
		return
	}
	//密码不返回给前端
	user.Password = ""
	token := jwtUtil.CreateJwt(user.Name)
	var UserDTO = UsersDTO{
		Name:  user.Name,
		Token: token,
	}
	context.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "登录成功",
		"data": UserDTO,
	})
}
func UserRegister(context *gin.Context) {
	var user = Users{}
	if err := context.ShouldBindJSON(&user); err != nil {
		context.JSON(200, gin.H{
			"code": 400,
			"msg":  "请求参数错误",
		})
		return
	}
	//判断账号长度不能少于3位
	if len(user.Name) > 3 {
		context.JSON(http.StatusOK, gin.H{
			"code": "2002",
			"msg":  "用户名不能少于3位",
		})
		return
	}
	//判断密码长度不能少于6位
	if len(user.Password) > 6 {
		context.JSON(http.StatusOK, gin.H{
			"code": "2001",
			"msg":  "密码长度不能少于6位",
		})
		return
	}
	message := db.Where("name = ?", user.Name).First(&user)
	if message.RowsAffected == 1 {
		//说明账号已经被注册
		context.JSON(http.StatusOK, gin.H{
			"code": 10001,
			"msg":  "账号已被注册",
		})
		return
	}
	//可以注册
	db.Create(&user)
	context.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "注册成功",
	})
	return

}
