package router

import (
	"tiktok_project/global"
	"tiktok_project/model"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

func InitBaseValidator() {
	initUserRegistValidator()
}

func initUserRegistValidator() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("is_unique", func(fl validator.FieldLevel) bool {
			if value, ok := fl.Field().Interface().(string); ok {
				var user model.User
				err := global.DB.Model(&user).Where("name = ?", value).First(&user).Error
				if err == gorm.ErrRecordNotFound {
					return true
				}
			}
			return false
		})
	}
}
