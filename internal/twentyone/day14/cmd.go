package day14

import (
	"bufio"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/withmandala/go-log"

	"github.com/cstaaben/adventofcode/internal/config"
)

func Cmd() *cobra.Command {
	d := new(dayFourteen)

	return &cobra.Command{
		Use:     "day_fourteen",
		Aliases: []string{"fourteen"},
		RunE:    d.runE,
	}
}

type dayFourteen struct {
	logger *log.Logger
}

func (d *dayFourteen) runE(_ *cobra.Command, _ []string) error {
	conf, err := config.New(2022, 14)
	if err != nil {
		return fmt.Errorf("parsing config: %w", err)
	}

	d.logger = log.New(os.Stderr)
	if conf.Debug {
		d.logger = d.logger.WithDebug()
	}
	d.logger.Debug("-----> Day Fourteen")

	var file *os.File
	file, err = os.Open(conf.InputFile)
	if err != nil {
		d.logger.Error("error opening input file:", err)
		return nil
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	runOne, runTwo := conf.ShouldRun()
	if runOne {
		d.partOne(scanner)
	}
	if runTwo {
		d.partTwo(scanner)
	}

	return nil
}

func (d *dayFourteen) partOne(scanner *bufio.Scanner) {
	d.logger.Debug("----------> Part One")
}

func (d *dayFourteen) partTwo(scanner *bufio.Scanner) {
	d.logger.Debug("----------> Part Two")
}