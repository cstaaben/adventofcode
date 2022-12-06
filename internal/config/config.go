package config

import "github.com/spf13/viper"

type Config struct {
	Debug      bool `mapstructure:"debug"`
	RunPartOne bool `mapstructure:"part_one"`
	RunPartTwo bool `mapstructure:"part_two"`
	UseInput   bool `mapstructure:"use_input"`
	UseSample  bool `mapstructure:"use_sample"`
}

func New() (*Config, error) {
	c := new(Config)
	err := viper.Unmarshal(c)
	return c, err
}
