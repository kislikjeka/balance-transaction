package config

import "github.com/spf13/viper"

func InitConfig(filename string) error {
	viper.AddConfigPath("configs")
	viper.SetConfigName(filename)
	return viper.ReadInConfig()
}
