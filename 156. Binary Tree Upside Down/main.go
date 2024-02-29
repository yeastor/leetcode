package main

func main() {

}

func upsideDownBinaryTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	nextLeft := root.Left
	nextRight := root.Right

	root.Left = nil
	root.Right = nil
	return rotate(root, nextLeft, nextRight)
}

func rotate(root *TreeNode, left *TreeNode, right *TreeNode) *TreeNode {

	if left == nil {
		return root
	}
	nextLeft := left.Left
	nextRight := left.Right

	left.Right = root
	left.Left = right

	if nextLeft == nil {
		return left
	}

	return rotate(left, nextLeft, nextRight)
}
