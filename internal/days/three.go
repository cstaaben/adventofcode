package days

import (
	"bufio"
	"fmt"
	"os"

	"github.com/spf13/cobra"

	log "github.com/cstaaben/go-log"

	"github.com/cstaaben/adventofcode/internal/config"
)

func NewThreeCommand() *cobra.Command {
	d := &dayThree{
		logger: log.New(os.Stderr).WithoutTimestamp().WithoutFile(),
	}

	cmd := &cobra.Command{
		Use: "three",
		Run: d.run,
	}

	cmd.Flags().Bool("part_one", false, "Run part one of the day's puzzle")
	cmd.Flags().Bool("part_two", false, "Run part two of the day's puzzle")
	cmd.Flags().Bool("use_sample", false, "Use sample input file for the day. Expected file is input/three/sample.txt")
	cmd.Flags().Bool("use_input", false, "Use puzzle input file. Expected file is at input/three/input.txt.")

	return cmd
}

type dayThree struct {
	logger *log.Logger
}

func (d *dayThree) run(_ *cobra.Command, _ []string) {
	conf, err := config.New()
	if err != nil {
		d.logger.Error("error parsing config:", err)
		return
	}

	if conf.Debug {
		d.logger = d.logger.WithDebug()
	}
	d.logger.Debug("Day Three")

	var (
		filepath string
		file     *os.File
	)
	if conf.UseSample {
		filepath = "input/three/sample.txt"
	} else if conf.UseInput {
		filepath = "input/three/input.txt"
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

func (d *dayThree) partOne(scanner *bufio.Scanner) {
	d.logger.Debug("----------> Part One")

	cumulativePri := 0
	for scanner.Scan() {
		line := scanner.Text()
		mid := len(line) / 2
		one, two := line[:mid], line[mid:]

		common := intersect(one, two)
		priority := toPriority(common)

		d.logger.Debugf("common: %s\tcode: %d\tpriority: %d\n", string(common), common, priority)
		cumulativePri += priority
	}

	fmt.Println(cumulativePri)
}

func intersect(one, two string) rune {
	a := make(map[rune]struct{})
	b := make(map[rune]struct{})

	for _, r := range one {
		a[r] = struct{}{}
	}
	for _, r := range two {
		b[r] = struct{}{}
	}

	for r := range a {
		if _, found := b[r]; found {
			return r
		}
	}

	return 0
}

func toPriority(r rune) int {
	if r >= 'a' && r <= 'z' {
		return int(r - 96)
	} else if r >= 'A' && r <= 'Z' {
		return int(r - 38)
	}

	return 0
}

func (d *dayThree) partTwo(scanner *bufio.Scanner) {
	d.logger.Debug("----------> Part Two")

	lineCount := 0
	cumulativePriority := 0
	lines := make([]string, 3)
	for scanner.Scan() {
		line := scanner.Text()
		lineCount++

		switch lineCount {
		case 1:
			lines[0] = line
		case 2:
			lines[1] = line
		case 3:
			lines[2] = line
		}

		if lineCount%3 != 0 {
			continue
		}
		lineCount = 0

		common := nIntersect(lines...)
		priority := toPriority(common)
		cumulativePriority += priority
	}

	fmt.Println(cumulativePriority)
}

func nIntersect(lists ...string) rune {
	maps := make([]map[rune]struct{}, len(lists))

	for i, list := range lists {
		m := make(map[rune]struct{})
		for _, r := range list {
			m[r] = struct{}{}
		}

		maps[i] = m
	}

	for r := range maps[0] {
		allPresent := true

		for _, m := range maps[1:] {
			if _, found := m[r]; !found {
				allPresent = false
				break
			}
		}

		if allPresent {
			return r
		}
	}

	return 0
}
