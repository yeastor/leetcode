package main

import (
	"fmt"
	"reflect"
	"testing"
)

type TestCase struct {
	Vals [][]int
	Want []int
}

func getTestCases() []*TestCase {

	return []*TestCase{
		{
			Vals: [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			Want: []int{1, 2, 3, 6, 9, 8, 7, 4, 5},
		},
		{
			Vals: [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}},
			Want: []int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7},
		},
		{
			Vals: [][]int{{1, 2, 3, 4}},
			Want: []int{1, 2, 3, 4},
		},
		{
			Vals: [][]int{{1}, {2}, {3}, {4}},
			Want: []int{1, 2, 3, 4},
		},
	}
}

func TestMaxArea(t *testing.T) {
	testCases := getTestCases()

	for _, tc := range testCases {
		fmt.Printf("Start %v \n", tc.Vals)
		res := spiralOrder(tc.Vals)
		if !reflect.DeepEqual(res, tc.Want) {
			t.Errorf("error res = %v, w = %v\n", res, tc.Want)
		}

	}
}
