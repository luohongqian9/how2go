package main

import (
	"ginlearn/api"
	"ginlearn/global"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitDB() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	global.DB = db
}

func main() {
	InitDB()

	r := gin.Default()
	r.Use(gin.Logger())

	v1 := r.Group("/v1")
	{
		v1.Get("/user/:id", api.GetUserHandler)
	}

	r.Run("0.0.0.0:8080")
}
