package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func postorderTraversal(root *TreeNode) []int {
	res := []int{}
	startRecurce(root, &res)

	return res
}

func startRecurce(root *TreeNode, res *[]int) {
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
	startRecurce(root, res)
}

func moveRight(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}
	startRecurce(root, res)
}
