package day8

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/withmandala/go-log"

	"github.com/cstaaben/adventofcode/internal/config"
)

func Cmd() *cobra.Command {
	d := new(dayEight)

	return &cobra.Command{
		Use:     "day_eight",
		Aliases: []string{"eight"},
		RunE:    d.runE,
	}
}

type dayEight struct {
	logger *log.Logger
}

func (d *dayEight) runE(_ *cobra.Command, _ []string) error {
	conf, err := config.New(2021, 8)
	if err != nil {
		return fmt.Errorf("parsing config: %w", err)
	}

	d.logger = log.New(os.Stderr)
	if conf.Debug {
		d.logger = d.logger.WithDebug()
	}
	d.logger.Debug("-----> Day Eight")

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

func (d *dayEight) partOne(scanner *bufio.Scanner) {
	d.logger.Debug("----------> Part One")

	readings, err := d.readFile(scanner)
	if err != nil {
		d.logger.Error("error reading file:", err)
		return
	}

	count := 0
	for _, reading := range readings {
		for _, digit := range reading.digits {
			switch len(digit) {
			case 2, 4, 3, 7:
				count++
			}
		}
	}

	fmt.Printf("Part One:\n\tUnique count instances: %d\n", count)
}

func (d *dayEight) partTwo(scanner *bufio.Scanner) {
	d.logger.Debug("----------> Part Two")

	readings, err := d.readFile(scanner)
	if err != nil {
		d.logger.Error("error reading file:", err)
		return
	}

	sum := 0
	for _, reading := range readings {
		mapping := reading.determineMapping()

		d.logger.Debug("reading:", reading)

		for i, segments := range reading.digits {
			digit := newDigit(segments, mapping)

			if len(digit) != len(segments) {
				d.logger.Errorf("dropped segments when converting to digit:\n\tsegments: %s\n\tdigit: %s", segments, digit.segments())
				return
			}

			var v int
			switch i {
			case 0:
				v = digit.toInt() * 1000
			case 1:
				v = digit.toInt() * 100
			case 2:
				v = digit.toInt() * 10
			case 3:
				v = digit.toInt()
			}

			d.logger.Debugf("segments: %s\tmapped: %s\ti: %d\tval: %d\n", segments, digit.segments(), i, v)

			sum += v
		}
	}

	fmt.Println("Part Two:", sum)
}

func (d *dayEight) readFile(scanner *bufio.Scanner) ([]reading, error) {
	readings := make([]reading, 0)

	for scanner.Scan() {
		s := strings.Split(scanner.Text(), "|")
		r := reading{
			signals: make([]string, 0),
			digits:  make([]string, 0),
		}

		for _, signal := range strings.Split(s[0], " ") {
			if signal == "" || signal == " " {
				continue
			}

			r.signals = append(r.signals, signal)
		}

		for _, num := range strings.Split(s[1], " ") {
			if num == "" || num == " " {
				continue
			}

			r.digits = append(r.digits, num)
		}

		if strings.Join(s, "|") != r.String() {
			return nil, fmt.Errorf("mangled the input somehow: %v", r)
		}

		readings = append(readings, r)
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("reading file: %w\n", err)
	}

	return readings, nil
}
