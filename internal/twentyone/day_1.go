package twentyone

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/withmandala/go-log"

	"github.com/cstaaben/adventofcode/internal/config"
)

var DayOneCmd = &cobra.Command{
	Use:  "one",
	RunE: func(_ *cobra.Command, _ []string) error { return day1() },
}

func init() {
	DayOneCmd.Flags().Bool("part-one", false, "Run part one")
	DayOneCmd.Flags().Bool("part-two", false, "Run part one")
	DayOneCmd.Flags().BoolP("all", "a", false, "Run all parts")
	cobra.CheckErr(viper.BindPFlags(DayOneCmd.Flags()))
}

func day1() error {
	conf, err := config.New()
	if err != nil {
		return err
	}

	logger := log.New(os.Stderr)
	if conf.Debug {
		logger = logger.WithDebug()
	}
	logger.Debug("starting Day 1")

	inputContent, err := ioutil.ReadFile("./input/day_1.txt")
	if err != nil {
		logger.Errorf("error reading input file: %v\n", err)
		return err
	}

	readings := strings.Split(string(inputContent), "\n")

	fmt.Print("Day 1: ")

	runOne, runTwo := shouldRun()

	if runOne {
		partOne(readings, logger)
	}
	if runTwo {
		partTwo(readings, logger)
	}

	return nil
}

func shouldRun() (bool, bool) {
	var (
		runOne bool
		runTwo bool
	)
	if viper.GetBool("all") {
		return true, true
	}

	if viper.GetBool("part-one") {
		runOne = true
	}

	if viper.GetBool("part-two") {
		runTwo = true
	}

	return runOne, runTwo
}

func partOne(readings []string, logger *log.Logger) {
	logger.Debug("-----> Part One")
	increasedCnt := 0

	for i := 1; i < len(readings)-1; i++ {
		//logger.Debugf("i: %5v\ti-1: %5v\tincreased: %v\n", readings[i], readings[i-1], readings[i] > readings[i-1])

		current, _ := strconv.Atoi(readings[i])
		prev, _ := strconv.Atoi(readings[i-1])

		if current > prev {
			increasedCnt++
		}
	}

	fmt.Printf("\tPart 1: %d\n", increasedCnt)
}

func partTwo(readings []string, logger *log.Logger) {
	logger.Debug("-----> Part Two")
	increasedCnt := 0

	windowSize := 3
	for i := 0; i < len(readings); i++ {
		currentStart := i
		currentEnd := i + windowSize
		prevStart := currentStart - 1
		prevEnd := prevStart + windowSize

		if currentEnd >= len(readings) || prevStart < 0 {
			continue
		}

		currentWindow := getWindowSum(currentStart, currentEnd, readings)
		prevWindow := getWindowSum(prevStart, prevEnd, readings)

		logger.Debugf("i: %d\n\tCurrent window  [%4d, %4d): %v = %4d\n\tPrevious window [%4d, %4d): %v = %4d\n\tincreased: %v\n",
			i,
			currentStart,
			currentEnd,
			readings[currentStart:currentEnd],
			currentWindow,
			prevStart,
			prevEnd,
			readings[prevStart:prevEnd],
			prevWindow,
			currentWindow > prevWindow,
		)

		if currentWindow > prevWindow {
			increasedCnt++
		}
	}

	fmt.Printf("\tPart 2: %d\n", increasedCnt)
}

func getWindowSum(start, end int, readings []string) int {
	sum := 0
	sub := readings[start:end]
	if len(sub) != 3 {
		fmt.Printf("slice size: %d\n", len(sub))
	}
	for _, s := range sub {
		n, _ := strconv.Atoi(s)
		sum += n
	}

	return sum
}
