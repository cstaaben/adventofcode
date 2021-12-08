package day11

import (
	"bufio"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/withmandala/go-log"

	"github.com/cstaaben/adventofcode/internal/config"
)

func Cmd() *cobra.Command {
	d := new(dayEleven)

	return &cobra.Command{
		Use:     "day_eleven",
		Aliases: []string{"eleven"},
		RunE:    d.runE,
	}
}

type dayEleven struct {
	logger *log.Logger
}

func (d *dayEleven) runE(_ *cobra.Command, _ []string) error {
	conf, err := config.New(2022, 11)
	if err != nil {
		return fmt.Errorf("parsing config: %w", err)
	}

	d.logger = log.New(os.Stderr)
	if conf.Debug {
		d.logger = d.logger.WithDebug()
	}
	d.logger.Debug("-----> Day Eleven")

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

func (d *dayEleven) partOne(scanner *bufio.Scanner) {
	d.logger.Debug("----------> Part One")
}

func (d *dayEleven) partTwo(scanner *bufio.Scanner) {
	d.logger.Debug("----------> Part Two")
}
