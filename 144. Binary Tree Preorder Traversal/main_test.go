package main

import (
	"testing"
)

type TestCase struct {
	Vals *TreeNode
	Want []int
}

func getSimpleTree() *TreeNode {
	return &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val:   4,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   5,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val:   6,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   7,
				Left:  nil,
				Right: nil,
			},
		},
	}
}

func getTestCases() []*TestCase {

	return []*TestCase{
		&TestCase{
			Vals: getSimpleTree(),
			Want: []int{1, 2, 4, 5, 3, 6, 7},
		},
	}
}

func TestPreorder(t *testing.T) {
	testCases := getTestCases()

	for _, tc := range testCases {
		res := preorderTraversal(tc.Vals)
		if len(res) != len(tc.Want) {
			t.Errorf("error len not equals ressArr = %v", res)
			return
		}
		for i, val := range res {
			if val != tc.Want[i] {
				t.Errorf("error res = %v, w = %v\n, ressArr = %v", val, tc.Want[i], res)
			}
		}

	}
}
