package config

import "github.com/spf13/viper"

type Config struct {
	Year      int    `mapstructure:"year"`
	Debug     bool   `mapstructure:"debug"`
	InputFile string `mapstructure:"input_file"`
	// FetchInput  bool   `mapstructure:"fetch_input"`
	RunPartOne  bool `mapstructure:"part_one"`
	RunPartTwo  bool `mapstructure:"part_two"`
	RunAllParts bool `mapstructure:"all"`
}

func New() (*Config, error) {
	c := new(Config)
	viper.SetConfigFile(".adventofcode.yaml")
	err := viper.Unmarshal(c)
	return c, err
}

func (c *Config) ShouldRun() (bool, bool) {
	if c.RunAllParts {
		return true, true
	}

	return c.RunPartOne, c.RunPartTwo
}
