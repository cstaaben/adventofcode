package convert

import "testing"

func TestNumberToWord(t *testing.T) {
    t.Parallel()

	testCases := []struct {
		name         string
		num          int
		expectedWord string
	}{
		{
			name:         "single digit",
			num:          5,
			expectedWord: "five",
		},
		{
			name:         "less than 20",
			num:          15,
			expectedWord: "fifteen",
		},
		{
			name:         "25",
			num:          25,
			expectedWord: "twenty-five",
		},
		{
			name:         "50",
			num:          50,
			expectedWord: "fifty",
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			if s := NumberToWord(tc.num); s != tc.expectedWord {
				t.Errorf("expected %s but got %s\n", tc.expectedWord, s)
			}
		})
	}
}

func TestWordToNumber(t *testing.T) {
    t.Parallel()

    testCases := []struct{
        name string
        word string
        expectedNum int
    }{
        {
            name: "less than 10",
            word: "three",
            expectedNum: 3,
        },
        {
            name: "less than 20",
            word: "eighteen",
            expectedNum: 18,
        },
        {
            name: "multiple of 10",
            word: "thirty",
            expectedNum: 30,
        },
        {
            name: "even number",
            word: "forty-two",
            expectedNum: 42,
        },
        {
            name: "multiple of 5",
            word: "seventy-five",
            expectedNum: 75,
        },
    }

    for _, tc := range testCases {
        tc := tc
        t.Run(tc.name, func(t *testing.T) {
            if n := WordToNumber(tc.word); n != tc.expectedNum {
                t.Errorf("expected %d but received %d\n", tc.expectedNum, n)
            }
        })
    }
}
