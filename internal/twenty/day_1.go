package twenty

import (
	"bufio"
	"fmt"
	"os"
	"sort"

	"github.com/spf13/cobra"
	"github.com/withmandala/go-log"

	"github.com/cstaaben/adventofcode/internal/config"
	"github.com/cstaaben/adventofcode/internal/convert"
)

const total = 2020

var day1Cmd = &cobra.Command{
	Use: "one",
	RunE: func(_ *cobra.Command, _ []string) error {
		d := new(dayOne)
		return d.day1()
	},
}

type dayOne struct{}

func (d *dayOne) day1() error {
	conf, err := config.New(2021, 0)
	if err != nil {
		return err
	}

	logger := log.New(os.Stderr)
	if conf.Debug {
		logger = logger.WithDebug()
	}
	logger.Debug("-----> Day One")

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

func (d *dayOne) partOne(scanner *bufio.Scanner, logger *log.Logger) {
	logger.Debug("----------> Part One")

	expenses, err := convert.ScanNewlineInts(scanner)
	if err != nil {
		logger.Error("error scanning lines into integers:", err)
		return
	}

	sort.Ints(expenses)

	for i := range expenses {
		target := total - expenses[i]
		if target <= 0 {
			continue
		}

		idx := sort.SearchInts(expenses, target)
		if idx < len(expenses) && expenses[idx]+expenses[i] == total {
			fmt.Printf("\tPart One: %d\n", expenses[i]*expenses[idx])
			return
		}
	}
}

func (d *dayOne) partTwo(scanner *bufio.Scanner, logger *log.Logger) {
	logger.Debug("----------> Part Two")

	expenses, err := convert.ScanNewlineInts(scanner)
	if err != nil {
		logger.Error("error scanning lines into integers:", err)
		return
	}

	sort.Ints(expenses)

	for i := range expenses {
		if expenses[i] > total {
			continue
		}

		for j := range expenses {
			if expenses[i]+expenses[j] > total || expenses[j] > total {
				continue
			}

			target := total - expenses[i] - expenses[j]
			idx := sort.SearchInts(expenses, target)
			if idx < len(expenses) && expenses[i]+expenses[j]+expenses[idx] == total {
				fmt.Printf("\tPart Two: %d\n", expenses[i]*expenses[j]*expenses[idx])
				return
			}
		}
	}
}
