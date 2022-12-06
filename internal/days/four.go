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

	cmd.Flags().Bool("part_one", false, "Run part one of the day's puzzle")
	cmd.Flags().Bool("part_two", false, "Run part two of the day's puzzle")
	cmd.Flags().Bool("use_sample", false, "Use sample input file for the day. Expected file is input/four/sample.txt")
	cmd.Flags().Bool("use_input", false, "Use puzzle input file. Expected file is at input/four/input.txt.")

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

	var (
		filepath string
		file     *os.File
	)
	if conf.UseSample {
		filepath = "input/four/sample.txt"
	} else if conf.UseInput {
		filepath = "input/four/input.txt"
	}
	file, err = os.Open(filepath)
	if err != nil {
		d.logger.Error("error opening input file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	if conf.RunPartOne {
		d.partOne(scanner)
	}
	if conf.RunPartTwo {
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
