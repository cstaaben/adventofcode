package twentyone

import (
	"bufio"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/withmandala/go-log"

	"github.com/cstaaben/adventofcode/internal/config"
)

var day5Cmd = &cobra.Command{
	Use:     "day_five",
	Aliases: []string{"five"},
	RunE: func(_ *cobra.Command, _ []string) error {
		d := new(dayFive)
		return d.day5()
	},
}

type dayFive struct{}

func (d *dayFive) day5() error {
	conf, err := config.New()
	if err != nil {
		return fmt.Errorf("parsing config: %w", err)
	}

	logger := log.New(os.Stderr)
	if conf.Debug {
		logger = logger.WithDebug()
	}
	logger.Debug("-----> Day Five")

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

func (d *dayFive) partOne(scanner *bufio.Scanner, logger *log.Logger) {
	logger.Debug("----------> Part One")
}

func (d *dayFive) partTwo(scanner *bufio.Scanner, logger *log.Logger) {
	logger.Debug("----------> Part Two")
}
