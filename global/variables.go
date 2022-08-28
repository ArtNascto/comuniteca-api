package global

import (
	"comuniteca.com.br/api/dtos"
	"gorm.io/gorm"
)

var (
	Config dtos.Config
	DB     *gorm.DB
)
