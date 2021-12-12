package day8

import (
	"log"
	"strings"
)

func newDigit(segments string, mapping map[rune]rune) digit {
	d := digit{}

	for _, r := range segments {
		log.Printf("enabled segment %s is mapped to %s", string(r), string(mapping[r]))
		d[mapping[r]] = struct{}{}
	}

	return d
}

type digit map[rune]struct{}

func (d digit) toInt() int {
	switch len(d) {
	case 2:
		return 1
	case 3:
		return 7
	case 4:
		return 4
	case 7:
		return 8
	case 5: // 2,3,5
		_, bFound := d['b']
		_, eFound := d['e']

		if bFound { // only 5 has segment b enabled here
			return 5
		}

		if eFound { // only 2 has segment e enabled here
			return 2
		}

		return 3 // no segment b or e must mean 3
	case 6: // 0,6,9
		_, dFound := d['d']
		_, eFound := d['e']

		if !dFound { // only 0 has segment d disabled here
			return 0
		}

		if eFound { // only 6 has segment e enabled here
			return 6
		}

		return 9 // segment d and no segment e must mean 9
	default:
		return -1
	}
}

func (d digit) segments() string {
	sb := new(strings.Builder)

	for s := range d {
		sb.WriteString(string(s))
	}

	return sb.String()
}
