package apiserver

import "github.com/spf13/viper"

type Config struct {
	BindAddr string `toml:"bind_addr"`
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
	}

}
