package configutil

import (
	"github.com/spf13/viper"
	"manso.live/backend/golang-service/pkg/constant"
	"strings"
)

// InitConfig 使用 viper，根据配置文件名及类型初始化配置项
func InitConfig(filename, configType string) error {
	viper.SetConfigType(configType)
	viper.SetConfigFile(filename)
	viper.AddConfigPath(constant.DefaultConfigHome)
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	return viper.ReadInConfig()
}

// InitConfigForEnv 以 env 文件类型来初始化处理指定的配置文件
func InitConfigForEnv(filename string) error {
	return InitConfig(filename, constant.DefaultConfigEnv)
}
