package main

import (
	"fmt"

	"github.com/leetCodeTasks/datastruct/binary_tree"
)

func main() {
	myTree := binary_tree.New(6)
	myTree.Add(2)
	myTree.Add(8)
	myTree.Add(0)
	myTree.Add(4)
	myTree.Add(7)
	myTree.Add(9)
	myTree.Add(3)
	myTree.Add(5)

	fmt.Println(maxDepth(myTree))
}

func maxDepth(root *binary_tree.TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + max(maxDepth(root.Left), maxDepth(root.Right))
}

func max(x int, y int) int {
	if x > y {
		return x
	}
	return y
}
