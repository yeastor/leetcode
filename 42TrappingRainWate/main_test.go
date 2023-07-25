package main

import "testing"

type TestCase struct {
	Vals []int
	Want int
}

func TestFirstMissingPositive(t *testing.T) {
	testCases := getTestCases()

	for _, tc := range testCases {
		res := trap(tc.Vals)
		if res != tc.Want {
			t.Errorf("error res = %v, w = %v\n", res, tc.Want)
		}

	}
}

func getTestCases() []*TestCase {

	return []*TestCase{
		{
			Vals: []int{7, 2, 6, 4, 5},
			Want: 5,
		},
		{
			Vals: []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1},
			Want: 6,
		},
		{
			Vals: []int{4, 2, 0, 3, 2, 5},
			Want: 9,
		},
		{
			Vals: []int{1, 1, 1, 1, 1},
			Want: 0,
		},
		{
			Vals: []int{1, 1, 0, 1, 1},
			Want: 1,
		},
		{
			Vals: []int{1, 2, 3, 4, 5},
			Want: 0,
		},
		{
			Vals: []int{5, 4, 3, 2, 1},
			Want: 0,
		},
		{
			Vals: []int{5, 4, 6, 2, 7, 6, 7},
			Want: 6,
		},
	}
}
