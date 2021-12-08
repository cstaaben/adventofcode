package day2

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
	"github.com/withmandala/go-log"

	"github.com/cstaaben/adventofcode/internal/config"
)

func Cmd() *cobra.Command {
	d := new(dayTwo)

	return &cobra.Command{
		Use: "two",
		RunE: func(_ *cobra.Command, _ []string) error {
			return d.day2()
		},
	}
}

type dayTwo struct{}

func (d *dayTwo) day2() error {
	conf, err := config.New(2021, 0)
	if err != nil {
		return err
	}

	logger := log.New(os.Stderr)
	if conf.Debug {
		logger = logger.WithDebug()
	}
	logger.Debug("-----> Day Two")

	var file *os.File
	file, err = os.Open(conf.InputFile)
	if err != nil {
		logger.Error("error opening input file:", err)
		return nil
	}
	scanner := bufio.NewScanner(file)

	depth := 0
	horizontal := 0
	aim := 0
	for scanner.Scan() {
		line := scanner.Text()
		a := strings.Split(line, " ")
		direction := a[0]
		distance, _ := strconv.Atoi(a[1])

		switch direction {
		case "down":
			aim += distance
		case "up":
			aim -= distance
		case "forward":
			horizontal += distance
			depth += aim * distance
		default:
			logger.Warn("unexpected direction:", direction)
		}
	}
	if err = scanner.Err(); err != nil {
		logger.Error("error scanning input file:", err)
		return nil
	}

	fmt.Printf("Day 2:\n\tPart One: %d\n", depth*horizontal)

	return nil
}
