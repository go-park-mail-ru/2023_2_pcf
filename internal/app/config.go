package server

import (
	"github.com/spf13/viper"
)

type Config struct {
	AuthBindAddr   string `toml:"auth_bind_addr"`
	BindAddr       string `toml:"bind_addr"`
	DataBase       string `toml:"database_url"`
	LogLevel       string `toml:"log_level"`
	Redis_addr     string `toml:"redis_addr"`
	Redis_password string `toml:"redis_password"`
	Redis_db       int    `toml:"redis_db"`
	File_path      string `toml:"file_path"`
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
		AuthBindAddr:   viper.GetString("auth_bind_addr"),
		BindAddr:       viper.GetString("bind_addr"),
		DataBase:       viper.GetString("database_url"),
		LogLevel:       viper.GetString("log_level"),
		Redis_addr:     viper.GetString("redis_addr"),
		Redis_password: viper.GetString("redis_password"),
		Redis_db:       viper.GetInt("redis_db"),
		File_path:      viper.GetString("file_path"),
	}

}
