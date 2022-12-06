package days

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"

	"github.com/spf13/cobra"

	log "github.com/cstaaben/go-log"

	"github.com/cstaaben/adventofcode/internal/config"
)

func NewOneCommand() *cobra.Command {
	d := new(dayOne)

	cmd := &cobra.Command{
		Use: "one",
		Run: d.run,
	}

	cmd.Flags().Bool("part_one", false, "Run part one of the day's puzzle")
	cmd.Flags().Bool("part_two", false, "Run part two of the day's puzzle")
	cmd.Flags().Bool("use_sample", false, "Use sample input file for the day. Expected file is input/one/sample.txt")
	cmd.Flags().Bool("use_input", false, "Use puzzle input file. Expected file is at input/one/input.txt.")

	return cmd
}

type dayOne struct {
	logger *log.Logger
}

func (d *dayOne) run(_ *cobra.Command, _ []string) {
	conf, err := config.New()
	if err != nil {
		d.logger.Error("error parsing config:", err)
		return
	}

	if conf.Debug {
		d.logger = d.logger.WithDebug()
	}
	d.logger.Debug("Day One")

	var (
		filepath string
		file     *os.File
	)
	if conf.UseSample {
		filepath = "input/one/sample.txt"
	} else if conf.UseInput {
		filepath = "input/one/input.txt"
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

func (d *dayOne) partOne(scanner *bufio.Scanner) {
	d.logger.Debug("----------> Part One")

	highestIdx := 0
	highestCount := 0
	i := 0
	count := 0
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			if count > highestCount {
				highestCount = count
				highestIdx = i
			}

			i++
			count = 0
			continue
		}

		c, _ := strconv.Atoi(line)
		count += c
	}

	fmt.Printf("Elf %d has the most calories with %d\n", highestIdx+1, highestCount)
}

/**
 * Incorrect guesses:
 *  - 137545 (too low)
 */
func (d *dayOne) partTwo(scanner *bufio.Scanner) {
	d.logger.Debug("----------> Part Two")

	highestCount := make([]int, 0)
	i := 0
	count := 0
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			if len(highestCount) < 3 { // auto-add until we get 3 counts
				highestCount = append(highestCount, count)
				sort.Ints(highestCount)
			} else if highestCount[0] < count {
				// if the lowest high-count is lower than our current count,
				// replace the bottom-most high count
				d.logger.Debugf("Current highs: %v\n\tReplacing %d with %d\n", highestCount, highestCount[0], count)
				highestCount[0] = count
				sort.Ints(highestCount)
			}

			i++
			count = 0
			continue
		}

		c, _ := strconv.Atoi(line)
		count += c
	}

	result := 0
	for _, val := range highestCount {
		result += val
	}

	fmt.Println(result)
}
