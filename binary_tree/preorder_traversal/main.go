package main

import tree "github.com/leetCodeTasks/datastruct/binary_tree"

func main() {

}

type Res struct {
	Res []int
}

func preorderTraversal(root *tree.TreeNode) []int {
	r := Res{}
	r.traversal(root)
	return r.Res
}

func (r *Res) traversal(node *tree.TreeNode) {
	if node != nil {
		r.Res = append(r.Res, node.Val)
		r.traversal(node.Left)
		r.traversal(node.Right)
	}
}
