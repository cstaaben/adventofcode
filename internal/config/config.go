package config

import "github.com/spf13/viper"

type Config struct {
	Debug     bool   `mapstructure:"debug"`
	InputFile string `mapstructure:"input_file"`
}

func New() (*Config, error) {
	c := new(Config)
	err := viper.Unmarshal(c)
	return c, err
}
