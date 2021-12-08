package day6

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/withmandala/go-log"

	"github.com/cstaaben/adventofcode/internal/config"
)

func Cmd() *cobra.Command {
	d := new(daySix)

	cmd := &cobra.Command{
		Use:     "day_six",
		Aliases: []string{"six"},
		RunE:    d.runE,
	}

	cmd.Flags().IntVar(&d.days, "days", 80, "Number of days to run simulation for")
	cobra.CheckErr(viper.BindPFlags(cmd.Flags()))

	return cmd
}

type daySix struct {
	days   int
	logger *log.Logger
}

func (d *daySix) runE(_ *cobra.Command, _ []string) error {
	conf, err := config.New(2021, 6)
	if err != nil {
		return fmt.Errorf("parsing config: %w", err)
	}

	d.logger = log.New(os.Stderr)
	if conf.Debug {
		d.logger = d.logger.WithDebug()
	}
	d.logger.Debug("-----> Day Six")

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

func (d *daySix) partOne(scanner *bufio.Scanner) {
	d.logger.Debug("----------> Part One")

	fish, err := d.readFile(scanner)
	if err != nil {
		d.logger.Error("error scanning input file:", err)
		return
	}

	fishBuckets := d.sortIntoBuckets(fish)

	d.logger.Debugf("running simulation for %d days with a concurrency of %d", d.days, len(fish))
	for day := 0; day < d.days; day++ {
		dayBuckets := make(map[int]int64)
		for bucket := 7; bucket >= 0; bucket-- {
			if bucket == 0 {
				dayBuckets[8] = fishBuckets[0]
				dayBuckets[6] += fishBuckets[0]
			}

			dayBuckets[bucket] = fishBuckets[bucket+1]
		}

		for bucket := range dayBuckets {
			fishBuckets[bucket] = dayBuckets[bucket]
		}
	}

	total := int64(0)
	for bucket := range fishBuckets {
		total += fishBuckets[bucket]
	}

	fmt.Println("Part One:", total)
}

func (d *daySix) partTwo(scanner *bufio.Scanner) {
	d.logger.Debug("----------> Part Two is just part one with more days")
}

func (d *daySix) readFile(scanner *bufio.Scanner) ([]int, error) {
	fish := make([]int, 0)

	for scanner.Scan() {
		line := scanner.Text()
		ages := strings.Split(line, ",")

		for _, a := range ages {
			age, _ := strconv.Atoi(a)
			fish = append(fish, age)
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("reading file: %w", err)
	}

	return fish, nil
}

func (d *daySix) sortIntoBuckets(fish []int) map[int]int64 {
	buckets := make(map[int]int64)
	for i := 0; i < 9; i++ {
		buckets[i] = 0
	}

	for _, age := range fish {
		buckets[age]++
	}

	return buckets
}
