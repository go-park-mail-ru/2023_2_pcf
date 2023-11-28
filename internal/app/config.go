package server

import (
	"github.com/spf13/viper"
)

type Config struct {
	AuthBindAddr   string `toml:"auth_bind_addr"`
	SelectBindAddr string `toml:"select_bind_addr"`
	BindAddr       string `toml:"bind_addr"`
	DataBase       string `toml:"database_url"`
	LogLevel       string `toml:"log_level"`
	Redis_addr     string `toml:"redis_addr"`
	Redis_password string `toml:"redis_password"`
	Redis_db       int    `toml:"redis_db"`
	Redis_db_ul    int    `tom:"redis_db_ul"`
	Redis_db_csrf  int    `tom:"redis_db_csrf"`
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
		SelectBindAddr: viper.GetString("select_bind_addr"),
		AuthBindAddr:   viper.GetString("auth_bind_addr"),
		BindAddr:       viper.GetString("bind_addr"),
		DataBase:       viper.GetString("database_url"),
		LogLevel:       viper.GetString("log_level"),
		Redis_addr:     viper.GetString("redis_addr"),
		Redis_password: viper.GetString("redis_password"),
		Redis_db:       viper.GetInt("redis_db"),
		Redis_db_ul:    viper.GetInt("redis_db_ul"),
		Redis_db_csrf:  viper.GetInt("redis_db_csrf"),
		File_path:      viper.GetString("file_path"),
	}

}
