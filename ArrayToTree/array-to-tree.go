package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func arrayToTeeNode(arr []int) *TreeNode {
	if len(arr) == 0 || arr[0] == -1 {
		return nil
	}
	root := makeList(arr, 0)
	depthArr := []*TreeNode{root}

	makeNode(arr, depthArr, 0)

	return root
}

func makeNode(arr []int, depthNodes []*TreeNode, curIndex int) {
	var nextDepthNodes []*TreeNode
	var nextIndex int
	lenArr := len(arr)
	nextIndex = curIndex
	for i := 0; i < len(depthNodes); i++ {
		// add left
		nextIndex++
		if nextIndex == lenArr {
			return
		}
		depthNodes[i].Left = makeList(arr, nextIndex)
		if depthNodes[i].Left != nil {
			nextDepthNodes = append(nextDepthNodes, depthNodes[i].Left)
		}

		//add right
		nextIndex++
		if nextIndex == lenArr {
			return
		}
		depthNodes[i].Right = makeList(arr, nextIndex)
		if depthNodes[i].Right != nil {
			nextDepthNodes = append(nextDepthNodes, depthNodes[i].Right)
		}
	}

	makeNode(arr, nextDepthNodes, nextIndex)

}

func makeList(arr []int, index int) *TreeNode {
	if len(arr) < index+1 {
		return nil
	}
	val := arr[index]
	if val == -1 {
		return nil
	}
	return &TreeNode{
		Val:   val,
		Left:  nil,
		Right: nil,
	}
}
