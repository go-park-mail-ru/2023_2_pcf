package server

import (
	"github.com/spf13/viper"
)

type Config struct {
	AuthBindAddr string `toml:"auth_bind_addr"`
	BindAddr     string `toml:"survey_bind_addr"`
	Db           string `toml:"survey_db"`
	LogLevel     string `toml:"log_level"`
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
		AuthBindAddr: viper.GetString("auth_bind_addr"),
		BindAddr:     viper.GetString("survey_bind_addr"),
		Db:           viper.GetString("survey_db"),
		LogLevel:     viper.GetString("log_level"),
	}

}
