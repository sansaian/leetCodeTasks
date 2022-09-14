package main

import (
	tree "github.com/leetCodeTasks/datastruct/binary_tree"
)

func isValidBST(root *tree.TreeNode) bool {
	if root == nil {
		return true
	}

	if valid(root.Left, root.Val, true) && valid(root.Right, root.Val, false) {
		return isValidBST(root.Left) && isValidBST(root.Right)
	} else {
		return false
	}
}

func valid(root *tree.TreeNode, nodeVal int, isLess bool) bool {
	if root == nil {
		return true
	}
	if isLess {
		if root.Val >= nodeVal {
			return false
		}
		return valid(root.Left, nodeVal, isLess) && valid(root.Right, nodeVal, isLess)
	} else {
		if root.Val <= nodeVal {
			return false
		}
		return valid(root.Left, nodeVal, isLess) && valid(root.Right, nodeVal, isLess)
	}

}
