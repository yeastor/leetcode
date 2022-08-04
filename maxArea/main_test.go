package main

import "testing"

type TestCase struct {
	height []int
	Want   int
}

func getTestCases() []*TestCase {

	return []*TestCase{
		{
			height: []int{6, 4, 3, 1, 4, 6, 99, 62, 1, 2, 6},
			Want:   62,
		},
		{
			height: []int{0, 2},
			Want:   0,
		},
		{
			height: []int{1, 8, 6, 2, 5, 4, 8, 3, 7},
			Want:   49,
		},
		{
			height: []int{1, 8, 6, 2, 5, 10, 8, 3, 7},
			Want:   49,
		},
		{
			height: []int{1, 1},
			Want:   1,
		},
		{
			height: []int{1, 1, 1},
			Want:   2,
		},
		{
			height: []int{2, 3, 4, 5, 18, 17, 6},
			Want:   17,
		},
		{
			height: []int{1, 3, 2, 5, 25, 24, 5},
			Want:   24,
		},
	}
}

func TestMaxArea(t *testing.T) {
	testCases := getTestCases()

	for _, tc := range testCases {
		res := maxArea(tc.height)

		if res != tc.Want {
			t.Errorf("error %v res = %v, w = %v\n", tc.height, res, tc.Want)
		}
	}
}
