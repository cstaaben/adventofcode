package config

import "github.com/spf13/viper"

type Config struct {
	Year  int  `mapstructure:"year"`
	Debug bool `mapstructure:"debug"`
	// FetchInput  bool   `mapstructure:"fetch_input"`
	RunPartOne bool `mapstructure:"part_one"`
	RunPartTwo bool `mapstructure:"part_two"`
	UseSample  bool `mapstructure:"use_sample"`
	UseInput   bool `mapstructure:"use_input"`
}

func New() (*Config, error) {
	c := new(Config)
	viper.SetConfigFile(".adventofcode.yaml")
	err := viper.Unmarshal(c)
	return c, err
}
