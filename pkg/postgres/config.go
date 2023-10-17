package postgres

import "github.com/spf13/viper"

type Config struct {
	DatabaseURL string `toml:"database_url"`
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
		DatabaseURL: viper.GetString("database_url"),
	}

}
