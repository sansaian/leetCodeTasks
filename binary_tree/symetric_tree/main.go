package main

import tree "github.com/leetCodeTasks/datastruct/binary_tree"

func main() {

}

func isSymmetric(root *tree.TreeNode) bool {
	if root == nil {
		return true
	}
	return visit(root.Left, root.Right)
}

func visit(root1 *tree.TreeNode, root2 *tree.TreeNode) bool {
	if root1 == nil && root2 == nil {
		return true
	}

	if root1 == nil || root2 == nil {
		return false
	}

	if root1.Val != root2.Val {
		return false
	}

	return visit(root1.Right, root2.Left) && visit(root1.Left, root2.Right)
}
