package main

import (
	"math"

	tree "github.com/leetCodeTasks/datastruct/binary_tree"
)

func main() {

}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST(root *tree.TreeNode) bool {
	return valid(root, math.MinInt, math.MaxInt)
}

func valid(root *tree.TreeNode, min int, max int) bool {
	if root == nil {
		return true
	}
	if root.Val <= min || root.Val >= max {
		return false
	}
	left := valid(root.Left, min, root.Val)
	right := valid(root.Right, root.Val, max)
	return left && right
}
