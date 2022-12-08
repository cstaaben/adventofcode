package days

import (
	"bufio"
	"fmt"
	"os"

	"github.com/spf13/cobra"

	log "github.com/cstaaben/go-log"

	"github.com/cstaaben/adventofcode/internal/config"
)

func NewSixCommand() *cobra.Command {
	d := &daySix{
		logger: log.New(os.Stderr).WithoutTimestamp().WithoutFile(),
	}

	cmd := &cobra.Command{
		Use: "six",
		Run: d.run,
	}

	cmd.Flags().Bool("part_one", false, "Run part one of the day's puzzle")
	cmd.Flags().Bool("part_two", false, "Run part two of the day's puzzle")
	cmd.Flags().Bool("use_sample", false, "Use sample input file for the day. Expected file is input/six/sample.txt")
	cmd.Flags().Bool("use_input", false, "Use puzzle input file. Expected file is at input/six/input.txt.")

	return cmd
}

type daySix struct {
	logger *log.Logger
}

func (d *daySix) run(_ *cobra.Command, _ []string) {
	conf, err := config.New()
	if err != nil {
		d.logger.Error("error parsing config:", err)
		return
	}

	if conf.Debug {
		d.logger = d.logger.WithDebug()
	}
	d.logger.Debug("Day Six")

	var (
		filepath string
		file     *os.File
	)
	if conf.UseSample {
		filepath = "input/six/sample.txt"
	} else if conf.UseInput {
		filepath = "input/six/input.txt"
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

func (d *daySix) partOne(scanner *bufio.Scanner) {
	d.logger.Debug("----------> Part One")

	for scanner.Scan() {
		line := scanner.Text()

		for start, end := 0, 3; end < len(line); start, end = start+1, end+1 {
			chars := make(map[string]struct{})
			for i := start; i <= end; i++ {
				if _, found := chars[string(line[i])]; found {
					break
				}

				chars[string(line[i])] = struct{}{}
			}

			if len(chars) == 4 {
				d.logger.Debugf("Index: %d\tCharacters: %v\n", end+1, chars)
				fmt.Println(end + 1)
				break
			}
		}
	}
}

func (d *daySix) partTwo(scanner *bufio.Scanner) {
	d.logger.Debug("----------> Part Two")

	for scanner.Scan() {
		line := scanner.Text()

		for start, end := 0, 13; end < len(line); start, end = start+1, end+1 {
			chars := make(map[string]struct{})
			for i := start; i <= end; i++ {
				if _, found := chars[string(line[i])]; found {
					break
				}

				chars[string(line[i])] = struct{}{}
			}

			if len(chars) == (end-start)+1 {
				d.logger.Debugf("Index: %d\tCharacters: %v\n", end+1, chars)
				fmt.Println(end + 1)
				break
			}
		}
	}
}
