package router

import (
	"zaplearn/internal/controller"
	"zaplearn/internal/middleware"

	"github.com/gin-gonic/gin"
)

func Register(r *gin.Engine) {
	r.Use(middleware.AccessLog())
	r.GET("/ping", controller.Ping)
}
