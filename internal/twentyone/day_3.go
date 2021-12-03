package twentyone

import (
	"bufio"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/withmandala/go-log"

	"github.com/cstaaben/adventofcode/internal/config"
)

var day3Cmd = &cobra.Command{
	Use:     "day_three",
	Aliases: []string{"three"},
	RunE: func(_ *cobra.Command, _ []string) error {
		d := new(dayThree)
		return d.day3()
	},
}

type dayThree struct{}

func (d *dayThree) day3() error {
	conf, err := config.New()
	if err != nil {
		return fmt.Errorf("parsing config: %w", err)
	}

	logger := log.New(os.Stderr)
	if conf.Debug {
		logger = logger.WithDebug()
	}
	logger.Debug("-----> Day Three")

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

func (d *dayThree) partOne(scanner *bufio.Scanner, logger *log.Logger) {
	logger.Debug("----------> Part One")
}

func (d *dayThree) partTwo(scanner *bufio.Scanner, logger *log.Logger) {
	logger.Debug("----------> Part Two")
}
