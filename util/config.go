package util

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	DBDriver      string `mapstructure:"DB_DRIVER"`
	DBSource      string `mapstructure:"DB_SOURCE"`
	ServerAddress string `mapstructure:"SERVER_ADDRESS"`
}

func LoadConfig(path string) (config Config, err error) {
	localEnv := viper.New()
	localEnv.AddConfigPath(path)
	localEnv.SetConfigName("app")
	localEnv.SetConfigType("env")

	localEnv.AutomaticEnv()

	err = localEnv.ReadInConfig()

	if err == nil {
		err = localEnv.Unmarshal(&config)
		return
	}

	if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
		return
	}

	err = nil

	sysEnv := viper.New()
	sysEnv.AutomaticEnv()
	sysEnv.BindEnv("DB_SOURCE")
	sysEnv.BindEnv("DB_DRIVER")
	sysEnv.BindEnv("SERVER_ADDRESS")

	err = sysEnv.Unmarshal(&config)

	log.Print(config)
	return

}
