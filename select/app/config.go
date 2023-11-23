package server

import (
	"github.com/spf13/viper"
)

type Config struct {
	BindAddr string `toml:"select_bind_addr"`
	LogLevel string `toml:"log_level"`
}

func Parse(configPath string) error {
	viper.SetConfigName("apiserver")
	viper.SetConfigType("toml")
	viper.AddConfigPath(configPath)

	err := viper.ReadInConfig()
	if err != nil {
		return err
	}

	return nil
}

func NewConfig() *Config {
	return &Config{
		BindAddr: viper.GetString("select_bind_addr"),
		LogLevel: viper.GetString("log_level"),
	}

}
