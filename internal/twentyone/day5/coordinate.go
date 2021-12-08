package day5

import (
	"fmt"
	"strconv"
)

type coordinate struct {
	x int
	y int
}

func (c coordinate) String() string {
	return fmt.Sprintf("%d,%d", c.x, c.y)
}

func (c coordinate) equals(that coordinate) bool {
	return c.y == that.y && c.x == that.x
}

func (c coordinate) applyDirection(d direction) coordinate {
	c.x += d.x
	c.y += d.y

	return c
}

func newCoordinateFromString(x, y string) (c coordinate, err error) {
	c.x, err = strconv.Atoi(x)
	if err != nil {
		return
	}

	c.y, err = strconv.Atoi(y)
	if err != nil {
		return
	}

	return
}
