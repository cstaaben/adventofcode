package convert

import (
	"bufio"
	"log"
	"strconv"
)

func ScanIntoInts(scanner *bufio.Scanner) ([]int, error) {
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
