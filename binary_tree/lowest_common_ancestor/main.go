package main

import "github.com/leetCodeTasks/datastruct/binary_tree"

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
}

func lowestCommonAncestor(root, p, q *binary_tree.TreeNode) *binary_tree.TreeNode {
	if root == nil {
		return root
	}
	if p.Val < root.Val && q.Val < root.Val {
		return lowestCommonAncestor(root.Left, p, q)
	}
	if p.Val > root.Val && q.Val > root.Val {
		return lowestCommonAncestor(root.Right, p, q)
	}
	if p.Val <= root.Val && q.Val >= root.Val {
		return root
	}
	return root
}
