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
			Vals: []int{},
			Want: nil,
		},
		{
			Vals: []int{1},
			Want: []int{1},
		},
		{
			Vals: []int{1, 2, 3, 4, 5},
			Want: []int{4, 5, 2, 3, 1},
		},
	}
}

func TestSay(t *testing.T) {
	testCases := getTestCases()

	for _, tc := range testCases {
		inTree := arrayToTeeNode(tc.Vals)
		resTree := upsideDownBinaryTree(inTree)
		resArr := preorderTraversal(resTree)
		if !reflect.DeepEqual(resArr, tc.Want) {
			t.Errorf("error res = %v, w = %v\n", resArr, tc.Want)
		}

	}
}
