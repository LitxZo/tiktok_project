package global

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	Logger *zap.SugaredLogger
	DB     *gorm.DB
)
