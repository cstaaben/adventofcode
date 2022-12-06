package days

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"

	log "github.com/cstaaben/go-log"

	"github.com/cstaaben/adventofcode/internal/config"
)

var (
	oppShapes = map[string]string{
		"A": "X", // rock
		"B": "Y", // paper
		"C": "Z", // scissors
	}
	scores = map[string]int{
		"X": 1,
		"Y": 2,
		"Z": 3,
	}
	losesTo = map[string]string{
		"X": "Y",
		"Y": "Z",
		"Z": "X",
	}
	beats = map[string]string{
		"X": "Z",
		"Y": "X",
		"Z": "Y",
	}

	drawPoints = 3
	winPoints  = 6
)

func NewTwoCommand() *cobra.Command {
	d := &dayTwo{
		logger: log.New(os.Stderr).WithoutTimestamp().WithoutFile(),
	}

	cmd := &cobra.Command{
		Use: "two",
		Run: d.run,
	}

	cmd.Flags().Bool("part_one", false, "Run part one of the day's puzzle")
	cmd.Flags().Bool("part_two", false, "Run part two of the day's puzzle")
	cmd.Flags().Bool("use_sample", false, "Use sample input file for the day. Expected file is input/two/sample.txt")
	cmd.Flags().Bool("use_input", false, "Use puzzle input file. Expected file is at input/two/input.txt.")

	return cmd
}

type dayTwo struct {
	logger *log.Logger
}

func (d *dayTwo) run(_ *cobra.Command, _ []string) {
	conf, err := config.New()
	if err != nil {
		d.logger.Error("error parsing config:", err)
		return
	}

	if conf.Debug {
		d.logger = d.logger.WithDebug()
	}
	d.logger.Debug("Day Two")

	var (
		filepath string
		file     *os.File
	)
	if conf.UseSample {
		filepath = "input/two/sample.txt"
	} else if conf.UseInput {
		filepath = "input/two/input.txt"
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

/**
 * Incorrect answers (reason):
 *  - 11663 (too low)
 */
func (d *dayTwo) partOne(scanner *bufio.Scanner) {
	d.logger.Debug("----------> Part One")

	score := 0
	for scanner.Scan() {
		moves := strings.Split(scanner.Text(), " ")
		opp, me := moves[0], moves[1]
		points := roundPoints(opp, me)

		d.logger.Debugf("opponent: %s\tme: %s\tscore: %d\n", opp, me, points)
		score += points
	}

	fmt.Println(score)
}

func roundPoints(opp, me string) (total int) {
	total += scores[me] // add score for shape
	if oppShapes[opp] == me {
		total += drawPoints
		return total
	}

	switch opp {
	case "A":
		if me == "Y" {
			total += winPoints
		}
	case "B":
		if me == "Z" {
			total += winPoints
		}
	case "C":
		if me == "X" {
			total += winPoints
		}
	default:
		panic(fmt.Sprintf("unexpected oppenent move value: %v", opp))
	}

	return total
}

func (d *dayTwo) partTwo(scanner *bufio.Scanner) {
	d.logger.Debug("----------> Part Two")
	d.logger.Debugf("Outcomes:\n\tX: lose\n\tY: draw\n\tZ: win\n")

	score := 0
	for scanner.Scan() {
		moves := strings.Split(scanner.Text(), " ")
		opp, me := moves[0], determineMyMove(moves[0], moves[1])
		points := roundPoints(opp, me)

		d.logger.Debugf("opponent: %s\toutcome: %s\tme: %s\tpoints: %d\n", opp, moves[1], me, points)
		score += points
	}

	fmt.Println(score)
}

// X = lose
// Y = draw
// Z = win
func determineMyMove(opp, outcome string) string {
	switch outcome {
	case "X":
		return beats[oppShapes[opp]]
	case "Y":
		return oppShapes[opp]
	case "Z":
		return losesTo[oppShapes[opp]]
	default:
		return ""
	}
}
