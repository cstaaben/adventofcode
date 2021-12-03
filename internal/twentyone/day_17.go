package twentyone

import (
	"bufio"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/withmandala/go-log"

	"github.com/cstaaben/adventofcode/internal/config"
)

var day17Cmd = &cobra.Command{
	Use:     "day_seventeen",
	Aliases: []string{"seventeen"},
	RunE: func(_ *cobra.Command, _ []string) error {
		d := new(daySeventeen)
		return d.day17()
	},
}

type daySeventeen struct{}

func (d *daySeventeen) day17() error {
	conf, err := config.New()
	if err != nil {
		return fmt.Errorf("parsing config: %w", err)
	}

	logger := log.New(os.Stderr)
	if conf.Debug {
		logger = logger.WithDebug()
	}
	logger.Debug("-----> Day Seventeen")

	var file *os.File
	file, err = os.Open(conf.InputFile)
	if err != nil {
		logger.Error("error opening input file:", err)
		return nil
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	runOne, runTwo := conf.ShouldRun()
	if runOne {
		d.partOne(scanner, logger)
	}
	if runTwo {
		d.partTwo(scanner, logger)
	}

	return nil
}

func (d *daySeventeen) partOne(scanner *bufio.Scanner, logger *log.Logger) {
	logger.Debug("----------> Part One")
}

func (d *daySeventeen) partTwo(scanner *bufio.Scanner, logger *log.Logger) {
	logger.Debug("----------> Part Two")
}
