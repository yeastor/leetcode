package main

import "sort"

type WeightNode struct {
	Val  int
	Next *WeightNode
}

func MakeNode(val int, prev *WeightNode) *WeightNode {
	node := &WeightNode{
		Val:  val,
		Next: nil,
	}

	if prev != nil {
		prev.Next = node
	}
	return node
}

func lastStoneWeight(stones []int) int {
	sort.Ints(stones)
	root := makeList(stones)
	return smashStones(root)
}

func inserNode(root *WeightNode, prevNode *WeightNode, toInsertNode *WeightNode) {

	if root.Next == nil {
		if root.Val >= toInsertNode.Val {
			root.Next = toInsertNode
		} else {
			oldroot := *root
			toInsertNode.Next = &oldroot
			*root = *toInsertNode
		}
		return
	}
	if root.Val <= toInsertNode.Val {

		if prevNode == nil {
			oldroot := *root
			toInsertNode.Next = &oldroot
			*root = *toInsertNode
			return
		}
		toInsertNode.Next = root
		prevNode.Next = toInsertNode
		toInsertNode.Next = root
		return
	}
	inserNode(root.Next, root, toInsertNode)

}

func smashStones(root *WeightNode) int {
	if root == nil {
		return 0
	}
	if root.Next == nil {
		return root.Val
	}
	smashRes := root.Val - root.Next.Val

	if root.Next != nil && smashRes > root.Next.Val {
		root.Next.Val = smashRes
		root = root.Next
		return smashStones(root)
	} else {
		root = root.Next.Next
	}

	if root == nil {
		return smashRes
	}

	if smashRes > 0 {
		toInsertNode := MakeNode(smashRes, nil)
		inserNode(root, nil, toInsertNode)
	}
	return smashStones(root)
}

func makeList(arr []int) *WeightNode {
	root := MakeNode(arr[len(arr)-1], nil)
	prevNode := root
	for i := len(arr) - 2; i >= 0; i-- {
		prevNode = MakeNode(arr[i], prevNode)
	}
	return root
}

func main() {

}
