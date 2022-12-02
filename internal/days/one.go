package days

import (
    "bufio"
    "os"

    "github.com/spf13/cobra"
    log "github.com/cstaaben/go-log"

    "github.com/cstaaben/adventofcode/internal/config"
    "strconv"
    "fmt"
    "sort"
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
func (d *dayOne) partTwo(scanner *bufio.Scanner, logger *log.Logger) {
    logger.Debug("----------> Part Two")

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
                logger.Debugf("Current highs: %v\n\tReplacing %d with %d\n", highestCount, highestCount[0], count)
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
