package main

import tree "github.com/leetCodeTasks/datastruct/binary_tree"

func main() {

}

type Res struct {
	Res []int
}

func inorderTraversal(root *tree.TreeNode) []int {
	r := Res{}
	r.dfs(root)
	return r.Res
}

func (r *Res) dfs(node *tree.TreeNode) {
	if node != nil {
		r.dfs(node.Left)
		r.Res = append(r.Res, node.Val)
		r.dfs(node.Right)
	}
}
