package controller

import (
	"net/http"
	"zaplearn/internal/service"

	"github.com/gin-gonic/gin"
)

func Ping(c *gin.Context) { // 应该使用 *gin.Context
	msg := service.Ping()

	c.JSON(http.StatusOK, gin.H{
		"message": msg,
	})
}
