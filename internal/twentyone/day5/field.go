package day5

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

func (f field) applyLine(l line) field {
	d := l.determineDirection()

	for current := l.start; !current.equals(l.end); current = current.applyDirection(d) {
		f[current.y][current.x]++
	}

	return f
}

func (f field) overlapsOverThreshold(threshold int) int {
	count := 0

	for i := range f {
		for j := range f[i] {
			if f[i][j] > threshold {
				count++
			}
		}
	}

	return count
}
