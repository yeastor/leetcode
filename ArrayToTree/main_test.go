package main

import (
	"reflect"
	"testing"
)

type TestCase struct {
	Vals []int
	Want []int
}

func getTestCases() []*TestCase {

	return []*TestCase{
		{
			Vals: []int{1, -1, 2, -1, 3, -1, 4, -1, 5},
			Want: []int{1, 2, 3, 4, 5},
		},
		{
			Vals: []int{1, 2, 3, 4, -1, 5, 6, 7, -1, -1, 8, 9, -1, 10, 11, -1, 12, 13, 14},
			Want: []int{1, 2, 4, 7, 10, 11, 3, 5, 8, 12, 6, 9, 13, 14},
		},
		{
			Vals: []int{1, 2, -1, 3, 4, -1, 5, -1, 6, 7, -1, -1, 8},
			Want: []int{1, 2, 3, 5, 7, 4, 6, 8},
		},
		{
			Vals: []int{1, 2, 3, -1, 4, -1, 5, 6, -1, -1, 7},
			Want: []int{1, 2, 4, 6, 3, 5, 7},
		},
		{
			Vals: []int{1, -1, 2, 3},
			Want: []int{1, 2, 3},
		},
		{
			Vals: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15},
			Want: []int{1, 2, 4, 8, 9, 5, 10, 11, 3, 6, 12, 13, 7, 14, 15},
		},
		{
			Vals: []int{1, 2, 3, 4, 5, 6, 7},
			Want: []int{1, 2, 4, 5, 3, 6, 7},
		},
		{
			Vals: []int{1, 2, 3, 4, 5},
			Want: []int{1, 2, 4, 5, 3},
		},
		{
			Vals: []int{1, 2, 3, -1, -1, 4, 5},
			Want: []int{1, 2, 3, 4, 5},
		},
	}
}

func TestArrayToTree(t *testing.T) {
	testCases := getTestCases()

	for _, tc := range testCases {
		resTree := arrayToTeeNode(tc.Vals)
		arr := preorderTraversal(resTree)
		if !reflect.DeepEqual(arr, tc.Want) {
			t.Errorf("error res = %v, w = %v\n", arr, tc.Want)
		}

	}
}
