package main

import (
	"fmt"
	"testing"
)

type TestCase struct {
	Vals []int
	Want int
}

func getTestCases() []*TestCase {

	return []*TestCase{
		{
			Vals: []int{7, 6, 7, 6, 9},
			Want: 3,
		},
		{
			Vals: []int{4, 6, 4, 10},
			Want: 4,
		},
		{
			Vals: []int{4, 6, 4, 10},
			Want: 4,
		},
		{
			Vals: []int{10, 45, 45, 555, 554, 3, 3, 5, 6, 7, 3, 3, 3},
			Want: 0,
		},
		{
			Vals: []int{3, 7, 2},
			Want: 2,
		},

		{
			Vals: []int{1},
			Want: 1,
		},
		{
			Vals: []int{1, 5},
			Want: 4,
		},

		{
			Vals: []int{2, 7, 4, 1, 8, 1},
			Want: 1,
		},
	}
}

func TestLastStoneWeight(t *testing.T) {
	testCases := getTestCases()

	for _, tc := range testCases {
		res := lastStoneWeight(tc.Vals)
		if res != tc.Want {
			t.Errorf("error %v res = %v, w = %v\n", tc, res, tc.Want)
		}
	}

}

func printList(root *WeightNode) {
	fmt.Printf("%v ", root.Val)
	if root.Next == nil {
		return
	}
	printList(root.Next)
}
