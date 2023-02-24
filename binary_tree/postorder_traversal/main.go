package main

import (
	"fmt"

	tree "github.com/leetCodeTasks/datastruct/binary_tree"
)

func main() {
	myTree := tree.New(6)
	myTree.Add(2)
	myTree.Add(8)
	myTree.Add(0)
	myTree.Add(4)
	myTree.Add(7)
	myTree.Add(9)
	myTree.Add(3)
	myTree.Add(5)

	fmt.Println(postorderTraversal(myTree))
}

type Res struct {
	Res []int
}

func postorderTraversal(root *tree.TreeNode) []int {
	r := Res{}
	r.traversal(root)
	return r.Res
}

func (r *Res) traversal(node *tree.TreeNode) {
	if node != nil {

		r.traversal(node.Left)
		r.traversal(node.Right)
		r.Res = append(r.Res, node.Val)
	}
}
