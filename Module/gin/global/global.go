package global

import "gorm.io/gorm"

var (
	DB     *gorm.DB
	Config *config.ServerConfig
)
