package twentyone

import (
	"bufio"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/withmandala/go-log"

	"github.com/cstaaben/adventofcode/internal/config"
)

var day20Cmd = &cobra.Command{
	Use:     "day_twenty",
	Aliases: []string{"twenty"},
	RunE: func(_ *cobra.Command, _ []string) error {
		d := new(dayTwenty)
		return d.day20()
	},
}

type dayTwenty struct{}

func (d *dayTwenty) day20() error {
	conf, err := config.New()
	if err != nil {
		return fmt.Errorf("parsing config: %w", err)
	}

	logger := log.New(os.Stderr)
	if conf.Debug {
		logger = logger.WithDebug()
	}
	logger.Debug("-----> Day Twenty")

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

func (d *dayTwenty) partOne(scanner *bufio.Scanner, logger *log.Logger) {
	logger.Debug("----------> Part One")
}

func (d *dayTwenty) partTwo(scanner *bufio.Scanner, logger *log.Logger) {
	logger.Debug("----------> Part Two")
}
