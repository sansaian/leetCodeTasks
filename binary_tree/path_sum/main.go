package main

import tree "github.com/leetCodeTasks/datastruct/binary_tree"

func main() {

}

var target = 0

func hasPathSum(root *tree.TreeNode, targetSum int) bool {
	target = targetSum
	return dfs(root, 0)
}

func dfs(root *tree.TreeNode, curr int) bool {
	if root == nil {
		return false
	}
	if root.Left == nil && root.Right == nil {
		return (root.Val + curr) == target
	}
	curr += root.Val
	left := dfs(root.Left, curr)
	right := dfs(root.Right, curr)
	return left || right
}
