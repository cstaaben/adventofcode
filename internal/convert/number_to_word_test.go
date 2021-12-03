package convert

import "testing"

func TestNumberToWord(t *testing.T) {
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
