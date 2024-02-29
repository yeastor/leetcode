package main

func startRecurcePostorder(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}
	moveLeft(root.Left, res)
	moveRight(root.Right, res)
	*res = append(*res, root.Val)
}

func moveLeft(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}
	startRecurcePostorder(root, res)
}

func moveRight(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}
	startRecurcePostorder(root, res)
}
