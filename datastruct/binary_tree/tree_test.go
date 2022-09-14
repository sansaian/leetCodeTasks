package binary_tree

import (
	"fmt"
	"testing"
)

func Test_BinaryTree(t *testing.T) {
	myTree := New(7)
	myTree.Add(9)
	myTree.Add(5)
	myTree.Add(13)
	fmt.Println(fmt.Sprintf("root %d, left %v, right %v", myTree.Val, myTree.Left, myTree.Right))
	myTree.Print()
	fmt.Println("search", myTree.Search(5))
	myTree.Add(4)
	myTree.Add(8)
	myTree.Print()
	fmt.Println("search", myTree.Search(5))

	myTree.Delete(4)
	myTree.Print()
}
