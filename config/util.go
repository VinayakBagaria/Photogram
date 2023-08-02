package config

import "github.com/spf13/viper"

func GetEnvString(key string) string {
	return viper.GetString(key)
}
