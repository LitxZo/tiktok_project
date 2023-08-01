package conf

import (
	"tiktok_project/model"
	"time"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(viper.GetString("DB.dsn")), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	sqlDB, _ := db.DB()
	sqlDB.SetConnMaxLifetime(time.Hour)
	sqlDB.SetMaxIdleConns(viper.GetInt("DB.maxIdleConns"))
	sqlDB.SetMaxOpenConns(viper.GetInt("DB.maxOpenConns"))

	db.AutoMigrate(&model.User{})
	return db, nil
}