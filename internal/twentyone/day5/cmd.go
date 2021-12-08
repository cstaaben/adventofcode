package day5

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
	d := new(dayFive)

	return &cobra.Command{
		Use:     "day_five",
		Aliases: []string{"five"},
		RunE:    d.runE,
	}
}

type dayFive struct {
	logger *log.Logger
}

func (d *dayFive) runE(_ *cobra.Command, _ []string) error {
	conf, err := config.New(2021, 5)
	if err != nil {
		return fmt.Errorf("parsing config: %w", err)
	}

	d.logger = log.New(os.Stderr)
	if conf.Debug {
		d.logger = d.logger.WithDebug()
	}
	d.logger.Debug("-----> Day Five")

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

/**
Incorrect attempts
- 1372
*/
func (d *dayFive) partOne(scanner *bufio.Scanner) {
	d.logger.Debug("----------> Part One")

	lines, maxX, maxY, err := d.readFile(scanner)
	if err != nil {
		return
	}

	f := newField(maxX, maxY)
	f = d.mapField(f, lines, false)

	overlaps := f.overlapsOverThreshold(2)

	fmt.Println("Part One:", overlaps)
}

func (d *dayFive) partTwo(scanner *bufio.Scanner) {
	d.logger.Debug("----------> Part Two")

	lines, maxX, maxY, err := d.readFile(scanner)
	if err != nil {
		return
	}

	f := newField(maxX, maxY)
	f = d.mapField(f, lines, true)

	overlaps := f.overlapsOverThreshold(2)

	fmt.Println("Part Two:", overlaps)
}

func (d *dayFive) readFile(scanner *bufio.Scanner) ([]line, int, int, error) {
	var (
		maxX int
		maxY int
		err  error

		lines = make([]line, 0)
	)

	for scanner.Scan() {
		rawCoords := strings.Split(scanner.Text(), " -> ")
		rawStarts := strings.Split(rawCoords[0], ",")
		rawEnds := strings.Split(rawCoords[1], ",")

		var start, end coordinate
		start, err = newCoordinateFromString(rawStarts[0], rawStarts[1])
		if err != nil {
			d.logger.Error("error getting start coordinates:", err)
			return nil, 0, 0, err
		}

		end, err = newCoordinateFromString(rawEnds[0], rawEnds[1])
		if err != nil {
			d.logger.Error("error getting end coordinates:", err)
			return nil, 0, 0, err
		}

		if start.x > maxX {
			maxX = start.x
		}
		if start.y > maxY {
			maxY = start.y
		}
		if end.x > maxX {
			maxX = end.x
		}
		if end.y > maxY {
			maxY = end.y
		}

		lines = append(lines, line{start: start, end: end})
	}
	if err = scanner.Err(); err != nil {
		d.logger.Error("error reading file:", err)
		return nil, 0, 0, err
	}

	return lines, maxX, maxY, nil
}

func (d *dayFive) mapField(f field, lines []line, considerDiagonals bool) field {
	f2 := f
	for _, l := range lines {
		if l.isDiagonal() && !considerDiagonals {
			d.logger.Debugf("line (%v) is diagonal, skipping", l)
			continue
		}
		d.logger.Debugf("applying line: %v (diagonal: %t)", l, l.isDiagonal())

		f2 = f2.applyLine(l)
	}

	return f2
}
