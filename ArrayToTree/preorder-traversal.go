package main

func preorderTraversal(root *TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}
	startRecurse(root, &res)
	return res
}

func move(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}
	startRecurse(root, res)
}

func startRecurse(root *TreeNode, res *[]int) {
	*res = append(*res, root.Val)
	move(root.Left, res)
	move(root.Right, res)
}
