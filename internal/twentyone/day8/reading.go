package day8

import (
	"fmt"
	"log"
	"sort"
	"strings"
)

type reading struct {
	signals []string
	digits  []string
}

// map enabled to actual segment
func (r reading) determineMapping() map[rune]rune {
	m := make(map[rune]rune)

	sort.Sort(signals(r.signals))

	for _, s := range r.signals {
		switch len(s) {
		case 2: // 1
			log.Printf("mapping %s to 1 (cf)\n", s)
			if _, found := m[rune(s[0])]; !found {
				m[rune(s[0])] = 'c'
			}
			if _, found := m[rune(s[1])]; !found {
				m[rune(s[1])] = 'f'
			}
		case 3: // 7
			log.Printf("mapping %s to 7 (acf)\n", s)
			if _, found := m[rune(s[0])]; !found {
				m[rune(s[0])] = 'a'
			}
			if _, found := m[rune(s[1])]; !found {
				m[rune(s[1])] = 'c'
			}
			if _, found := m[rune(s[2])]; !found {
				m[rune(s[2])] = 'f'
			}
		case 4: // 4
			log.Printf("mapping %s to 4 (bcdf)\n", s)
			if _, found := m[rune(s[0])]; !found {
				m[rune(s[0])] = 'b'
			}
			if _, found := m[rune(s[1])]; !found {
				m[rune(s[1])] = 'c'
			}
			if _, found := m[rune(s[2])]; !found {
				m[rune(s[2])] = 'd'
			}
			if _, found := m[rune(s[3])]; !found {
				m[rune(s[3])] = 'f'
			}
		case 7: // 8
			log.Printf("mapping %s to 8 (abcdefg)\n", s)
			if _, found := m[rune(s[0])]; !found {
				m[rune(s[0])] = 'a'
			}
			if _, found := m[rune(s[1])]; !found {
				m[rune(s[1])] = 'b'
			}
			if _, found := m[rune(s[2])]; !found {
				m[rune(s[2])] = 'c'
			}
			if _, found := m[rune(s[3])]; !found {
				m[rune(s[3])] = 'd'
			}
			if _, found := m[rune(s[4])]; !found {
				m[rune(s[4])] = 'e'
			}
			if _, found := m[rune(s[5])]; !found {
				m[rune(s[5])] = 'f'
			}
			if _, found := m[rune(s[6])]; !found {
				m[rune(s[6])] = 'g'
			}
		}
	}

	if len(m) != 7 {
		log.Println("missing some segment(s):", m)
	}

	segMap := make(map[string]string)
	for e, a := range m {
		segMap[string(e)] = string(a)
	}
	log.Printf("segment mapping: %v\n", segMap)

	segments := make(map[string]string)
	for enabled, actual := range m {
		if mappings, found := segments[string(actual)]; found {
			segments[string(actual)] = mappings + string(enabled)
		} else {
			segments[string(actual)] = string(enabled)
		}
	}

	if len(segments) > 0 {
		for actualSegment, enabledSegments := range segments {
			if len(enabledSegments) > 1 {
				log.Printf("found duplicate mappings for segment %s: %s\n", actualSegment, enabledSegments)
			}
		}
	}

	return m
}

func (r reading) String() string {
	return fmt.Sprintf("%s | %s", strings.Join(r.signals, " "), strings.Join(r.digits, " "))
}

var _ sort.Interface = signals{}

type signals []string

func (s signals) Len() int {
	return len(s)
}

func (s signals) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s signals) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
