package main

import (
	"bytes"
	"strconv"
)

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func tree2str(root *TreeNode) string {

	var buffer bytes.Buffer

	print(root, &buffer)
	return string(buffer.Bytes()[1:])
}

func print(leaf *TreeNode, buffer *bytes.Buffer) {
	buffer.WriteString("(" + strconv.Itoa(leaf.Val))
	if leaf.Left != nil {
		print(leaf.Left, buffer)
		buffer.WriteString(")")
	}
	if leaf.Right != nil {
		print(leaf.Right, buffer)
		buffer.WriteString(")")
	}
	//if leaf.Right == nil {
	//	buffer.WriteString(")")
	//}
	//if leaf.Left == nil {
	//	buffer.WriteString(")")
	//}
	return
}
