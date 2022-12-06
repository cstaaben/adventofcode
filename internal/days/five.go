package days

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/spf13/cobra"

	log "github.com/cstaaben/go-log"

	"github.com/cstaaben/adventofcode/internal/config"
)

var (
	directionRegex = regexp.MustCompile(`^move (\d+) from (\d+) to (\d+)$`)
	stacksRegex    = regexp.MustCompile(`\d`)
)

func NewFiveCommand() *cobra.Command {
	d := &dayFive{
		logger: log.New(os.Stderr).WithoutTimestamp().WithoutFile(),
	}

	cmd := &cobra.Command{
		Use: "five",
		Run: d.run,
	}

	cmd.Flags().Bool("part_one", false, "Run part one of the day's puzzle")
	cmd.Flags().Bool("part_two", false, "Run part two of the day's puzzle")
	cmd.Flags().Bool("use_sample", false, "Use sample input file for the day. Expected file is input/five/sample.txt")
	cmd.Flags().Bool("use_input", false, "Use puzzle input file. Expected file is at input/five/input.txt.")

	return cmd
}

type dayFive struct {
	logger *log.Logger
}

func (d *dayFive) run(_ *cobra.Command, _ []string) {
	conf, err := config.New()
	if err != nil {
		d.logger.Error("error parsing config:", err)
		return
	}

	if conf.Debug {
		d.logger = d.logger.WithDebug()
	}
	d.logger.Debug("Day Five")

	var (
		filepath string
		file     *os.File
	)
	if conf.UseSample {
		filepath = "input/five/sample.txt"
	} else if conf.UseInput {
		filepath = "input/five/input.txt"
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

func (d *dayFive) partOne(scanner *bufio.Scanner) {
	d.logger.Debug("----------> Part One")

	crateMover9000 := func(count, sourceIdx, destIdx int, stacks []*list.List) {
		for i := 0; i < count; i++ {
			s := stacks[sourceIdx].Remove(stacks[sourceIdx].Back())
			d.logger.Debug("Moving", s)
			stacks[destIdx].PushBack(s)
		}
	}

	stacks := d.loadCrates(scanner)
	d.moveCrates(scanner, stacks, crateMover9000)

	loadMsg(stacks)
}

func (d *dayFive) loadCrates(scanner *bufio.Scanner) []*list.List {
	d.logger.Debug("Loading crate layout")

	stacks := make([]*list.List, 0)
	crateLines := make([]string, 0)
	lineCount := 0
	for scanner.Scan() {
		lineCount++
		line := scanner.Text()
		if len(line) == 0 {
			break
		} else if !strings.Contains(line, "[") {
			matches := stacksRegex.FindAllStringSubmatch(line, -1)
			for i := 0; i < len(matches); i++ {
				stacks = append(stacks, list.New())
			}
		} else {
			crateLines = append(crateLines, line)
		}
	}

	for i := len(crateLines) - 1; i >= 0; i-- {
		line := crateLines[i]
		curStack := 0
		for i := 1; i < len(line); i, curStack = i+4, curStack+1 {
			if string(line[i]) == " " || string(line[i]) == "" {
				continue
			}
			d.logger.Debugf("Adding %s to stack %d\n", string(line[i]), curStack+1)
			stacks[curStack].PushBack(string(line[i]))
		}
	}

	return stacks
}

func (d *dayFive) moveCrates(scanner *bufio.Scanner, stacks []*list.List, moveFn func(count, sourceIdx, destIdx int,
	stacks []*list.List)) {
	d.logger.Debug("Moving crates")
	for scanner.Scan() {
		line := scanner.Text()
		matches := directionRegex.FindStringSubmatch(line)
		if len(matches) < 4 {
			d.logger.Errorf("unexpected no matches: %s\n\t"+
				"\tMatches: %v\n", line, matches)
		}
		directions := matches[1:]

		count, _ := strconv.Atoi(directions[0])
		sourceIdx, _ := strconv.Atoi(directions[1])
		destIdx, _ := strconv.Atoi(directions[2])

		sourceIdx--
		destIdx--

		d.logger.Debugf("Count: %d\tFrom: %d\tTo: %d\n", count, sourceIdx, destIdx)

		moveFn(count, sourceIdx, destIdx, stacks)

		d.debugStacks(stacks)
	}
}

func loadMsg(stacks []*list.List) {
	var msg string
	for _, stack := range stacks {
		msg += stack.Back().Value.(string)
	}

	fmt.Println(msg)
}

func (d *dayFive) debugStacks(stacks []*list.List) {
	msg := "Stacks (bottom -> top):\n"
	for i, stack := range stacks {
		var crates string
		for el := stack.Front(); el != nil; el = el.Next() {
			crates += el.Value.(string)
		}

		msg += fmt.Sprintf("\t%d: %s\n", i+1, crates)
	}

	d.logger.Debug(msg)
}

func (d *dayFive) partTwo(scanner *bufio.Scanner) {
	d.logger.Debug("----------> Part Two")

	crateMover9001 := func(count, sourceIdx, destIdx int, stacks []*list.List) {
		carriedStack := list.New()
		src, dest := stacks[sourceIdx], stacks[destIdx]
		for i := 0; i < count; i++ {
			carriedStack.PushFront(stacks[sourceIdx].Remove(src.Back()))
		}

		var crates string
		for el := carriedStack.Front(); el != nil; el = el.Next() {
			crates += el.Value.(string)
		}

		d.logger.Debugf("Moving stack: %v\n", crates)
		dest.PushBackList(carriedStack)
		carriedStack.Init()
	}

	// load crates into stacks (in-memory)
	stacks := d.loadCrates(scanner)
	d.debugStacks(stacks)
	// follow directions and move them
	d.moveCrates(scanner, stacks, crateMover9001)
	// print end message
	loadMsg(stacks)
}
