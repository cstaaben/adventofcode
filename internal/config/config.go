package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Debug       bool   `mapstructure:"debug"`
	InputFile   string `mapstructure:"input_file"`
	RunPartOne  bool   `mapstructure:"part_one"`
	RunPartTwo  bool   `mapstructure:"part_two"`
	RunAllParts bool   `mapstructure:"all"`
}

func New(year, day int) (*Config, error) {
	c := new(Config)

	viper.SetDefault("input_file", fmt.Sprintf("input/%d/day_%d.txt", year, day))

	err := viper.Unmarshal(c)
	return c, err
}

func (c *Config) ShouldRun() (bool, bool) {
	if c.RunAllParts {
		return true, true
	}

	return c.RunPartOne, c.RunPartTwo
}
