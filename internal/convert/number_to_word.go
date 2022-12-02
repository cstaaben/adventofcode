package convert

import (
    "fmt"
    "strings"
)

var numberToWord = map[int]string{
	1:  "one",
	2:  "two",
	3:  "three",
	4:  "four",
	5:  "five",
	6:  "six",
	7:  "seven",
	8:  "eight",
	9:  "nine",
	10: "ten",
	11: "eleven",
	12: "twelve",
	13: "thirteen",
	14: "fourteen",
	15: "fifteen",
	16: "sixteen",
	17: "seventeen",
	18: "eighteen",
	19: "nineteen",
	20: "twenty",
	30: "thirty",
	40: "forty",
	50: "fifty",
	60: "sixty",
	70: "seventy",
	80: "eighty",
	90: "ninety",
}

func NumberToWord(n int) string {
	if s, found := numberToWord[n]; found {
		return s
	}

	r := n % 10
	return fmt.Sprintf("%s-%s", numberToWord[n-r], numberToWord[r])
}

func WordToNumber(s string) int {
    // no hyphen should mean it's in the map somewhere
    if !strings.Contains(s, "-") {
        return findValFromWord(s)
    }

    words := strings.Split(s, "-")
    return findValFromWord(words[0]) + findValFromWord(words[1])
}

func findValFromWord(s string) int {
    for n, word := range numberToWord {
        if strings.EqualFold(s, word) {
            return n
        }
    }

    return 0
}
