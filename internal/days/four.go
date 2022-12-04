package days

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/cobra"

	log "github.com/cstaaben/go-log"

	"github.com/cstaaben/adventofcode/internal/config"
)

func NewFourCommand() *cobra.Command {
	d := &dayFour{
		logger: log.New(os.Stderr).WithoutTimestamp().WithoutFile(),
	}

	cmd := &cobra.Command{
		Use: "four",
		Run: d.run,
	}

	cmd.Flags().StringP("input_file", "i", "", "Input file for puzzle")
	cmd.Flags().Bool("part_one", false, "Run part one of the day's puzzle")
	cmd.Flags().Bool("part_two", false, "Run part two of the day's puzzle")
	cmd.Flags().BoolP("all", "a", false, "Run all parts of the day's puzzle")

	return cmd
}

type dayFour struct {
	logger *log.Logger
}

func (d *dayFour) run(_ *cobra.Command, _ []string) {
	conf, err := config.New()
	if err != nil {
		d.logger.Error("error parsing config:", err)
		return
	}

	if conf.Debug {
		d.logger = d.logger.WithDebug()
	}
	d.logger.Debug("Day Four")

	var file *os.File
	file, err = os.Open(conf.InputFile)
	if err != nil {
		d.logger.Error("error opening input file:", err)
		return
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
}

func (d *dayFour) partOne(scanner *bufio.Scanner) {
	d.logger.Debug("----------> Part One")

	count := 0
	for scanner.Scan() {
		sections := strings.Split(scanner.Text(), ",")

		nums := strings.Split(sections[0], "-")
		a, _ := strconv.Atoi(nums[0])
		b, _ := strconv.Atoi(nums[1])

		nums = strings.Split(sections[1], "-")
		x, _ := strconv.Atoi(nums[0])
		y, _ := strconv.Atoi(nums[1])

		if surrounded(a, b, x, y) {
			count++
		}
	}

	fmt.Println(count)
}

func surrounded(a, b, c, d int) bool {
	return (a <= c && b >= d) || (c <= a && d >= b)
}

func (d *dayFour) partTwo(scanner *bufio.Scanner) {
	d.logger.Debug("----------> Part Two")

	count := 0
	for scanner.Scan() {
		sections := strings.Split(scanner.Text(), ",")

		nums := strings.Split(sections[0], "-")
		a, _ := strconv.Atoi(nums[0])
		b, _ := strconv.Atoi(nums[1])

		nums = strings.Split(sections[1], "-")
		x, _ := strconv.Atoi(nums[0])
		y, _ := strconv.Atoi(nums[1])

		if overlap(a, b, x, y) {
			count++
		}
	}

	fmt.Println(count)
}

func overlap(a, b, c, d int) bool {
	return surrounded(a, b, c, d) || ((a <= c && b <= d && b >= c) || (a >= c && a <= d && b >= d) || (c <= a && d <= b && d >= a) || (c >= a && c <= b && d >= b))
}
