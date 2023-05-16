package main

import (
	"testing"
)

type TestCase struct {
	Vals []int
	Want int
}

func getTestCases() []*TestCase {

	return []*TestCase{
		{
			Vals: []int{-2, 1, -3, 4, -1, 2, 1, -5, 4},
			Want: 6,
		},
		{
			Vals: []int{1},
			Want: 1,
		},
		{
			Vals: []int{5, 4, -1, 7, 8},
			Want: 23,
		},
		{
			Vals: []int{-2, 1},
			Want: 1,
		},
	}
}

func TestMaxArea(t *testing.T) {
	testCases := getTestCases()

	for _, tc := range testCases {
		res := maxSubArray(tc.Vals)
		if res != tc.Want {
			t.Errorf("error res = %v, w = %v\n", res, tc.Want)
		}

	}
}
