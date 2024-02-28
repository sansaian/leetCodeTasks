package main

import (
	tree "github.com/leetCodeTasks/datastruct/binary_tree"
)

func isValidBSTNotOptimal(root *tree.TreeNode) bool {
	if root == nil {
		return true
	}

	if validNotOptimal(root.Left, root.Val, true) && validNotOptimal(root.Right, root.Val, false) {
		return isValidBST(root.Left) && isValidBST(root.Right)
	} else {
		return false
	}
}

func validNotOptimal(root *tree.TreeNode, nodeVal int, isLess bool) bool {
	if root == nil {
		return true
	}
	if isLess {
		if root.Val >= nodeVal {
			return false
		}
		return validNotOptimal(root.Left, nodeVal, isLess) && validNotOptimal(root.Right, nodeVal, isLess)
	} else {
		if root.Val <= nodeVal {
			return false
		}
		return validNotOptimal(root.Left, nodeVal, isLess) && validNotOptimal(root.Right, nodeVal, isLess)
	}

}
