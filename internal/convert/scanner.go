package convert

import (
	"bufio"
	"log"
	"strconv"
	"strings"
)

func ScanNewlineInts(scanner *bufio.Scanner) ([]int, error) {
	res := make([]int, 0)
	for scanner.Scan() {
		i, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return nil, err
		}

		if i < 0 {
			log.Println(i)
		}

		res = append(res, i)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return res, nil
}

func ScanCommaInts(scanner *bufio.Scanner) ([]int, error) {
	res := make([]int, 0)

	for scanner.Scan() {
		ints := strings.Split(scanner.Text(), ",")
		for _, s := range ints {
			i, _ := strconv.Atoi(s)
			res = append(res, i)
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return res, nil
}
