package twentyone

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"

	"github.com/spf13/cobra"
	"github.com/withmandala/go-log"

	"github.com/cstaaben/adventofcode/internal/config"
	"github.com/cstaaben/adventofcode/internal/convert"
)

var day3Cmd = &cobra.Command{
	Use:     "day_three",
	Aliases: []string{"three"},
	RunE: func(_ *cobra.Command, _ []string) error {
		d := new(dayThree)
		return d.day3()
	},
}

type dayThree struct {
	logger *log.Logger
}

func (d *dayThree) day3() error {
	conf, err := config.New(2021, 3)
	if err != nil {
		return fmt.Errorf("parsing config: %w", err)
	}

	d.logger = log.New(os.Stderr)
	if conf.Debug {
		d.logger = d.logger.WithDebug()
	}
	d.logger.Debug("-----> Day Three")

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

func (d *dayThree) partOne(scanner *bufio.Scanner) {
	d.logger.Debug("----------> Part One")

	var (
		bits      []int
		lineCount int

		initOnce = new(sync.Once)
	)

	for scanner.Scan() {
		line := scanner.Text()
		lineCount++

		initOnce.Do(func() {
			bits = make([]int, len(line))
		})

		for i := range line {
			bit, _ := strconv.Atoi(string(line[i]))
			bits[i] += bit
		}
	}
	if err := scanner.Err(); err != nil {
		d.logger.Error("error reading file:", err)
		return
	}

	gamma := make([]int, len(bits))
	epsilon := make([]int, len(bits))
	for i := range bits {
		if bits[i] > lineCount/2 {
			gamma[i] = 1
			epsilon[i] = 0
		} else {
			gamma[i] = 0
			epsilon[i] = 1
		}
	}

	g, err := convert.BinaryToDecimal(gamma)
	if err != nil {
		d.logger.Error("error converting gamma binary:", err)
		return
	}

	d.logger.Debugf("gamma: %v = %d", binaryNum(gamma), g)

	var e int64
	e, err = convert.BinaryToDecimal(epsilon)
	if err != nil {
		d.logger.Error("error converting epsilon binary:", err)
		return
	}

	d.logger.Debugf("epsilon: %v = %d", binaryNum(epsilon), e)

	fmt.Printf("Part One: %d\n", e*g)
}

type binaryNum []int

func (bn binaryNum) String() string {
	sb := new(strings.Builder)
	for i := range bn {
		sb.WriteString(strconv.Itoa(bn[i]))
	}
	return sb.String()
}

/**
Incorrect attempts:
3312583
*/
func (d *dayThree) partTwo(scanner *bufio.Scanner) {
	d.logger.Debug("----------> Part Two")

	// read in all numbers
	bins := make([]binaryNum, 0)
	for scanner.Scan() {
		line := scanner.Text()
		b := binaryNum(make([]int, len(line)))

		for i := range line {
			b[i], _ = strconv.Atoi(string(line[i]))
		}

		bins = append(bins, b)
	}

	co2KeepFn := func(candidates []binaryNum, i int) (keepOnes, keepZeroes bool) {
		ones, zeroes := d.countBits(candidates, i)

		keepOnes = ones < zeroes
		keepZeroes = zeroes <= ones

		return
	}

	o2KeepFn := func(candidates []binaryNum, i int) (keepOnes, keepZeroes bool) {
		ones, zeroes := d.countBits(candidates, i)

		keepOnes = ones >= zeroes
		keepZeroes = zeroes > ones

		return
	}

	d.logger.Debug("searching for O2 candidate")
	o2Candidate := d.findRating(bins, o2KeepFn)

	d.logger.Debug("searching for CO2 candidate")
	co2Candidate := d.findRating(bins, co2KeepFn)

	o2Rating, _ := convert.BinaryToDecimal(o2Candidate)
	co2Rating, _ := convert.BinaryToDecimal(co2Candidate)

	fmt.Printf("Part Two:\n\tCO2 scrubber reading: %s = %d\n\tO2 generator reading: %s = %d\n\tLife support rating: %d\n",
		co2Candidate,
		co2Rating,
		o2Candidate,
		o2Rating,
		co2Rating*o2Rating,
	)
}

func (d *dayThree) findRating(candidates []binaryNum, keepFn func([]binaryNum, int) (bool, bool)) binaryNum {
	bitCount := len(candidates[0])
	for i := 0; i < bitCount; i++ {
		keepOnes, keepZeroes := keepFn(candidates, i)
		candidates = d.pruneCandidates(candidates, i, keepOnes, keepZeroes)

		d.logger.Debugf("candidates after checking bit index %d: %v\n", i, candidates)
		if len(candidates) == 1 {
			return candidates[0]
		}
	}

	return nil
}

func (d *dayThree) countBits(candidates []binaryNum, bitIndex int) (ones, zeroes int) {
	for j := range candidates {
		switch candidates[j][bitIndex] {
		case 1:
			ones++
		case 0:
			zeroes++
		}
	}

	return
}

func (d *dayThree) pruneCandidates(candidates []binaryNum, bitIndex int, keepOne, keepZero bool) []binaryNum {
	c := make([]binaryNum, 0)

	for j := range candidates {
		switch candidates[j][bitIndex] {
		case 1:
			if keepOne {
				c = append(c, candidates[j])
			}
		case 0:
			if keepZero {
				c = append(c, candidates[j])
			}
		}
	}

	return c
}
