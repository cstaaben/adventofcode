package day5

import (
	"fmt"
)

type line struct {
	start coordinate
	end   coordinate
}

func (l line) determineDirection() direction {
	d := direction{}

	switch {
	case l.start.x < l.end.x:
		d.x = 1
	case l.start.x > l.end.x:
		d.x = -1
	case l.start.x == l.end.x:
		d.x = 0
	}

	switch {
	case l.start.y < l.end.y:
		d.y = 1
	case l.start.y > l.end.y:
		d.y = -1
	case l.start.y == l.end.y:
		d.y = 0
	}

	return d
}

func (l line) isDiagonal() bool {
	horizontal := l.start.x != l.end.x && l.start.y == l.end.y
	vertical := l.start.x == l.end.x && l.start.y != l.end.y
	return !(vertical || horizontal)
}

func (l line) String() string {
	return fmt.Sprintf("%v -> %v", l.start, l.end)
}
