package service

import "zaplearn/internal/global"

func Ping() string {
	global.Logger.Info("ping service called")
	return "pong"
}
