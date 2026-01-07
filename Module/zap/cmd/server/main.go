package main

import (
	"zaplearn/internal/config"
	"zaplearn/internal/global"
	"zaplearn/internal/router"

	"github.com/gin-gonic/gin"
)

func main() {
	logger, err := config.InitLogger()
	if err != nil {
		panic(err)
	}

	global.Logger = logger
	// global.Logger.Sync() - 刷新日志缓冲区。 zap 日志库会先将日志写入缓冲区， Sync() 方法会将缓冲区中的日志强制写入到目标（如文件、标准输出等）
	defer global.Logger.Sync()

	global.Logger.Info("server start")
	// gin.Default() - 创建一个带默认中间件（Logger 和 Recovery）的 Gin 引擎
	// gin.New() - 创建一个不带任何中间件的 Gin 引擎，更灵活，需要手动添加所需中间件
	r := gin.New()

	router.Register(r)

	r.Run("0.0.0.0:8080")
}
