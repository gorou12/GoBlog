package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// ここからGin
	router := gin.Default()
	router.LoadHTMLGlob("content/*.html")
	router.GET("/", index)
	router.GET("/admin", admin)
	router.POST("/post", post)

	router.Run()
}

func index(c *gin.Context) {
	c.HTML(200, "index.html", gin.H{"text": "hogehoge"})
}

func admin(c *gin.Context) {
	// ここからGORM
	dsn := "user:Password01@tcp(127.0.0.1:3306)/goblog"
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	c.HTML(200, "admin.html", gin.H{})
}

func post(c *gin.Context) {
	title := c.PostForm("title")
	content := c.PostForm("content")
	c.String(200, "Title: %s, Content: %s", title, content)
}
