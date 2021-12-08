package day4

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/withmandala/go-log"

	"github.com/cstaaben/adventofcode/internal/config"
)

func Cmd() *cobra.Command {
	d := new(dayFour)
	return &cobra.Command{
		Use:     "day_four",
		Aliases: []string{"four"},
		RunE: func(_ *cobra.Command, _ []string) error {
			return d.day4()
		},
	}
}

type dayFour struct {
	logger *log.Logger
}

func (d *dayFour) day4() error {
	conf, err := config.New(2021, 4)
	if err != nil {
		return fmt.Errorf("parsing config: %w", err)
	}

	d.logger = log.New(os.Stderr)
	if conf.Debug {
		d.logger = d.logger.WithDebug()
	}
	d.logger.Debug("-----> Day Four")

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

func (d *dayFour) partOne(scanner *bufio.Scanner) {
	d.logger.Debug("----------> Part One")

	boards, drawOrder, err := d.readFile(scanner)
	if err != nil {
		d.logger.Error("error reading file:", err)
		return
	}

	for i, b := range boards {
		d.logger.Debugf("board %d:\n%v\n", i, b)
	}

	var (
		foundWinner   bool
		winningBoard  board
		winningNumber int
	)

	for _, draw := range drawOrder {
		d.logger.Debugf("drew number: %d\n", draw)

		for i, b := range boards {
			b = b.applyDraw(draw)
			boards[i] = b

			d.logger.Debugf("\nboard %d after draw:\n%s", i, b)

			won := b.hasWon()
			if won {
				d.logger.Debug("found winner at:", i)
				winningNumber = draw
				winningBoard = b
				foundWinner = true
				break
			}
		}

		if foundWinner {
			break
		}
	}

	score := winningBoard.calculateScore()
	fmt.Printf("Part One: %d * %d = %d\n", score, winningNumber, score*winningNumber)
}

func (d *dayFour) partTwo(scanner *bufio.Scanner) {
	d.logger.Debug("----------> Part Two")

	boards, drawOrder, err := d.readFile(scanner)
	if err != nil {
		d.logger.Error("error reading file:", err)
		return
	}

	for i, b := range boards {
		d.logger.Debugf("board %d:\n%v\n", i, b)
	}

	var (
		mostRecentWinner  board
		mostRecentWinDraw int
		winnerCount       int
	)

	for _, draw := range drawOrder {
		d.logger.Debugf("drew number: %d\n", draw)

		for i, b := range boards {
			if won := b.hasWon(); won {
				d.logger.Debugf("skipping board %d that has already won", i)
				continue
			}

			b = b.applyDraw(draw)
			boards[i] = b

			d.logger.Debugf("\nboard %d after draw:\n%s", i, b)

			won := b.hasWon()
			if won {
				d.logger.Debug("found winner at:", i)

				mostRecentWinDraw = draw
				mostRecentWinner = b
				winnerCount++

				if winnerCount == len(boards) {
					break
				}
			}
		}

		if winnerCount == len(boards) {
			break
		}
	}

	score := mostRecentWinner.calculateScore()
	fmt.Printf("Part Two: %d * %d = %d\n", score, mostRecentWinDraw, score*mostRecentWinDraw)
}

func (d *dayFour) readFile(scanner *bufio.Scanner) ([]board, []int, error) {
	var (
		drawOrder []int
		lines     []string
	)

	d.logger.Debug("reading file")

	for scanner.Scan() {
		if lines == nil {
			l := scanner.Text()
			nums := strings.Split(l, ",")
			drawOrder = make([]int, len(nums))

			for i := range nums {
				drawOrder[i], _ = strconv.Atoi(nums[i])
			}

			lines = make([]string, 0)
			continue
		}

		l := scanner.Text()
		if len(l) == 0 {
			continue
		}

		lines = append(lines, l)
	}

	d.logger.Debugf("number of draws: %d", len(drawOrder))

	boardCount := len(lines) / 5
	boards := make([]board, boardCount)

	d.logger.Debug("number of boards:", boardCount)

	for i := 0; i < len(lines); i++ {
		boardNum := i / 5
		row := i % 5

		if row == 0 { // only create a new board when we reach a new board
			boards[boardNum] = board{}
		}

		sAra := strings.Split(lines[i], " ")
		skips := 0
		boards[boardNum][row] = [5]*entry{}
		for j, s := range sAra {
			if len(s) == 0 {
				skips++
				continue
			}

			e := new(entry)
			e.n, _ = strconv.Atoi(s)
			boards[boardNum][row][j-skips] = e
		}
	}

	return boards, drawOrder, nil
}

var match = color.New(color.FgGreen).SprintFunc()

type entry struct {
	n       int
	matched bool
}

func (e *entry) String() string {
	if e.matched {
		return match(fmt.Sprintf("%2d", e.n))
	}

	return fmt.Sprintf("%2d", e.n)
}

type board [5][5]*entry

func (b board) applyDraw(n int) board {
	for i := range b {
		for j := range b[i] {
			if b[i][j].n == n {
				b[i][j].matched = true
			}
		}
	}

	return b
}

func (b board) hasWon() bool {
	for i := 0; i < 5; i++ {
		rowMatch := true
		columnMatch := true
		for j := range b[i] {
			if !b[i][j].matched {
				rowMatch = false
			}

			if !b[j][i].matched {
				columnMatch = false
			}
		}

		if rowMatch || columnMatch {
			return true
		}
	}

	return false
}

func (b board) String() string {
	sb := new(strings.Builder)

	for i := range b {
		row := make([]string, len(b[i]))
		for j := range b[i] {
			row[j] = b[i][j].String()
		}

		sb.WriteString(strings.Join(row, " ") + "\n")
	}

	return sb.String()
}

func (b board) calculateScore() int {
	score := 0
	for i := range b {
		for j := range b[i] {
			if !b[i][j].matched {
				score += b[i][j].n
			}
		}
	}

	return score
}
