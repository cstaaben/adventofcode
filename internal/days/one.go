package days

import (
	"bufio"
	"os"

	"github.com/spf13/cobra"
	log "github.com/withmandala/go-log"

	"github.com/cstaaben/adventofcode/internal/config"
)

func NewOneCommand() *cobra.Command {
	d := new(dayOne)

	cmd := &cobra.Command{
		Use: "one",
		Run: d.run,
	}

	cmd.Flags().StringP("input_file", "i", "", "Input file for puzzle")
	cmd.Flags().Bool("part_one", false, "Run part one of the day's puzzle")
	cmd.Flags().Bool("part_two", false, "Run part two of the day's puzzle")
	cmd.Flags().BoolP("all", "a", false, "Run all parts of the day's puzzle")

	return cmd
}

type dayOne struct{}

func (d *dayOne) run(_ *cobra.Command, _ []string) {
	logger := log.New(os.Stderr)

	conf, err := config.New()
	if err != nil {
		logger.Error("error parsing config:", err)
		return
	}

	if conf.Debug {
		logger = logger.WithDebug()
	}
	logger.Debug("-----> Day One")

	var file *os.File
	file, err = os.Open(conf.InputFile)
	if err != nil {
		logger.Error("error opening input file:", err)
		return
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
}

func (d *dayOne) partOne(scanner *bufio.Scanner, logger *log.Logger) {
	logger.Debug("----------> Part One")
}

func (d *dayOne) partTwo(scanner *bufio.Scanner, logger *log.Logger) {
	logger.Debug("----------> Part Two")
}
