package conf

import (
	"fmt"

	"github.com/spf13/viper"
)

func InitConfig() {
	viper.SetConfigName("settings")
	viper.SetConfigType("yml")
	viper.AddConfigPath("./conf/")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Sprintf("Load COnfig Error: %s", err.Error()))
	}

	fmt.Println("Init Config Success")
}

const Islike = 0 // 点赞的状态

const Unlike = 1 //取消点赞状态

const LikeAction = 1 // 点赞的行为

const DefaultRedisValue = -1 // 默认值
