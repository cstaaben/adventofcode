package day7

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"

	"github.com/spf13/cobra"
	"github.com/withmandala/go-log"

	"github.com/cstaaben/adventofcode/internal/config"
	"github.com/cstaaben/adventofcode/internal/convert"
)

func Cmd() *cobra.Command {
	d := new(daySeven)

	return &cobra.Command{
		Use:     "day_seven",
		Aliases: []string{"seven"},
		RunE:    d.runE,
	}
}

type daySeven struct {
	logger *log.Logger
}

func (d *daySeven) runE(_ *cobra.Command, _ []string) error {
	conf, err := config.New(2021, 7)
	if err != nil {
		return fmt.Errorf("parsing config: %w", err)
	}

	d.logger = log.New(os.Stderr)
	if conf.Debug {
		d.logger = d.logger.WithDebug()
	}
	d.logger.Debug("-----> Day Seven")

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

func (d *daySeven) partOne(scanner *bufio.Scanner) {
	d.logger.Debug("----------> Part One")

	positions, err := convert.ScanCommaInts(scanner)
	if err != nil {
		d.logger.Error("error reading file:", err)
		return
	}

	median := positions[len(positions)/2]
	cost := int64(0)
	for _, pos := range positions {
		distance := int64(math.Abs(float64(pos - median)))
		cost += distance
	}

	fmt.Printf("Part One:\n\tMedian: %d\n\tCost: %d\n", median, cost)
}

func (d *daySeven) partTwo(scanner *bufio.Scanner) {
	d.logger.Debug("----------> Part Two")

	positions, err := convert.ScanCommaInts(scanner)
	if err != nil {
		d.logger.Error("error reading file:", err)
		return
	}
	sort.Ints(positions)

	sum := sumInts(positions)
	avg := sum / len(positions)

	cost := int64(0)
	for _, pos := range positions {
		dist := int64(math.Abs(float64(pos - avg)))

		for i := int64(1); i <= dist; i++ {
			cost += i
		}
	}

	fmt.Printf("Part Two:\n\tCost: %d\n", cost)
}

func sumInts(nums []int) (sum int) {
	for _, n := range nums {
		sum += n
	}

	return
}
