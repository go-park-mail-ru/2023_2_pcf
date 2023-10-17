package server

import (
	"AdHub/internal/app/store"

	"github.com/spf13/viper"
)

type Config struct {
	BindAddr string `toml:"bind_addr"`
	Store    *store.Config
}

func Parse(configPath string) error {
	viper.SetConfigName("apiserver")
	viper.SetConfigType("toml")
	viper.AddConfigPath(configPath)

	err := viper.ReadInConfig()
	if err != nil { // Handle errors reading the config file
		return err
	}

	return nil
}

func NewConfig() *Config {
	return &Config{
		BindAddr: viper.GetString("bind_addr"),
		Store:    store.NewConfig(),
	}

}
