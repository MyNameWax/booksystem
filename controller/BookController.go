package controller

import (
	"github.com/gin-gonic/gin"
)

type Books struct {
	Id              int
	BookName        string
	BookPrice       string
	BookAuthor      string
	BookDescription string
	Status          int
}

func BookList(context *gin.Context) {
	var books []Books
	var RealBooks []Books
	db.Find(&books)
	for i := 0; i < len(books); i++ {
		if books[i].Status == 1 {
			RealBooks = append(RealBooks, books[i])
		}
	}
	context.JSON(200, gin.H{
		"code":    200,
		"message": "successful",
		"data":    RealBooks,
	})
}
func BookDetail(context *gin.Context) {
	// 1. 接收参数
	bookName := context.PostForm("name")
	// 2. 根据参数查询
	var books = Books{}
	result := db.Where("book_name =?", bookName).First(&books)
	if result.RowsAffected == 0 {
		context.JSON(200, gin.H{
			"code":    3001,
			"message": "书籍不存在",
		})
		return
	}
	// 3. 脱敏 返回数据
	books.Id = 00000
	context.JSON(200, gin.H{
		"code":    200,
		"message": "successful",
		"data":    books,
	})
}
func BookDeleted(context *gin.Context) {
	// 1.获取需要删除的书籍名称
	bookName := context.PostForm("name")
	// 2.查询数据库 图书是否存在
	var books = Books{}
	result := db.Where("book_name = ?", bookName).First(&books)
	if result.RowsAffected == 0 {
		// 3.不存在 抛异常
		context.JSON(200, gin.H{
			"code":    3001,
			"message": "图书不存在",
		})
		return
	}
	// 3.存在 删除
	resultDeleted := db.Where("book_name = ?", bookName).Delete(&books)
	if resultDeleted.RowsAffected != 0 {
		// 5. 返回
		context.JSON(200, gin.H{
			"code":    200,
			"message": "删除成功",
		})
	}
}
func BookAdd(context *gin.Context) {
	//TODO 测试哦哦哦
}
