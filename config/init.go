package config

import (
	"strings"

	"github.com/spf13/viper"
)

func Init(name, path string) error {
	// name of the config file
	viper.SetConfigName(name)

	// path to look for the config file in
	viper.AddConfigPath(path)

	err := viper.ReadInConfig()
	if err != nil {
		return err
	}

	// define replacer
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)

	viper.AutomaticEnv()

	return nil
}

func GetConfigValue(key string) string {
	return viper.GetString(key)
}
