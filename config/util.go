package config

import "github.com/spf13/viper"

func GetConfigValue(key string) string {
	return viper.GetString(key)
}
