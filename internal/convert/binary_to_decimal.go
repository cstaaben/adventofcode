package convert

import (
	"errors"
	"math"
)

// BinaryToDecimal takes an array of integers representing a binary number and converts it to its decimal representation
func BinaryToDecimal(bin []int) (int64, error) {
	n := int64(0)
	for i := range bin {
		if bin[i] < 0 || bin[i] > 1 {
			return 0, errors.New("unexpected value in binary array")
		} else if bin[i] == 0 {
			continue
		}

		pow := float64(len(bin) - (i + 1))
		b := int64(math.Pow(2.0, pow))

		n += b
	}

	return n, nil
}
