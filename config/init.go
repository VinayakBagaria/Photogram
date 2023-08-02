package config

import (
	"log"
	"strings"

	"github.com/spf13/viper"
)

func Init() error {
	// name of the config file
	viper.SetConfigName("config")

	// path to look for the config file in
	viper.AddConfigPath("./config")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalln("Unable to read the config file: %w", err)
	}

	// define replacer
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)

	viper.AutomaticEnv()

	return nil
}
