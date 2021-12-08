package day5

import "testing"

func TestLine_isDiagonal(t *testing.T) {
	testCases := []struct {
		name             string
		l                line
		expectedDiagonal bool
	}{
		{
			name: "diagonal",
			l: line{
				start: coordinate{
					x: 6,
					y: 4,
				},
				end: coordinate{
					x: 2,
					y: 0,
				},
			},
			expectedDiagonal: true,
		},
		{
			name: "horizontal",
			l: line{
				start: coordinate{
					x: 0,
					y: 9,
				},
				end: coordinate{
					x: 5,
					y: 9,
				},
			},
			expectedDiagonal: false,
		},
		{
			name: "vertical",
			l: line{
				start: coordinate{
					x: 0,
					y: 0,
				},
				end: coordinate{
					x: 0,
					y: 9,
				},
			},
			expectedDiagonal: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			diagonal := tc.l.isDiagonal()
			if diagonal != tc.expectedDiagonal {
				t.Errorf("actual diagonal value (%t) did not match expected diagonal value (%t)", diagonal, tc.expectedDiagonal)
			}
		})
	}
}
