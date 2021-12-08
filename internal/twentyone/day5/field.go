package day5

import (
	"strconv"
	"strings"
)

type field [][]int

func newField(cols, rows int) (f field) {
	c := cols + 1
	r := rows + 1

	f = make([][]int, r)

	for i := 0; i < r; i++ {
		f[i] = make([]int, c)
	}

	return
}

func (f field) String() string {
	sb := new(strings.Builder)

	for i := range f {
		for j := range f[i] {
			var s string
			switch f[i][j] {
			case 0:
				s = "."
			default:
				s = strconv.Itoa(f[i][j])
			}

			sb.WriteString(s)
		}

		sb.WriteString("\n")
	}

	return sb.String()
}

func (f field) applyLine(l line) field {
	d := l.determineDirection()
	current := l.start

	for { // do-while
		f[current.y][current.x]++

		if current.equals(l.end) {
			break
		}

		current = current.applyDirection(d)
	}

	return f
}

func (f field) overlapsOverThreshold(threshold int) int {
	count := 0

	for i := range f {
		for j := range f[i] {
			if f[i][j] >= threshold {
				count++
			}
		}
	}

	return count
}
