package main

import "fmt"

type BinaryTree struct {
	Root *Node
}

type Node struct {
	Left  *Node
	Right *Node
	Value int
}

func New() *BinaryTree {
	return &BinaryTree{
		Root: nil,
	}
}

func (b *BinaryTree) Delete(val int) {
	delete(b.Root, nil, val)
}

func delete(node *Node, upNode *Node, val int) {
	if node == nil {
		return
	}

	switch {
	case node.Value > val:
		delete(node.Left, node, val)
	case node.Value < val:
		delete(node.Right, node, val)
	case node.Value == val:
		if node.Value < upNode.Value {
			upNode.Left = nil
			return
		}
		upNode.Right = nil
		return
	}
	return
}

func (b *BinaryTree) Add(val int) {
	b.Root = add(b.Root, val)
}

func add(leaf *Node, val int) *Node {
	if leaf == nil {
		return &Node{
			Left:  nil,
			Right: nil,
			Value: val,
		}
	}
	switch {
	case leaf.Value > val:
		leaf.Left = add(leaf.Left, val)
	case leaf.Value < val:
		leaf.Right = add(leaf.Right, val)
	case leaf.Value == val:
		return leaf
	}
	return leaf
}

func (b *BinaryTree) Search(val int) *Node {
	return search(b.Root, val)
}

func search(leaf *Node, val int) *Node {
	if leaf.Value == val {
		return leaf
	}
	switch {
	case leaf.Value > val:
		return search(leaf.Left, val)
	case leaf.Value < val:
		return search(leaf.Right, val)
	}
	return nil
}

func (b *BinaryTree) Print() {
	print(b.Root)
}

func print(leaf *Node) {
	fmt.Println(leaf.Value)
	if leaf.Right != nil {
		print(leaf.Right)
	}
	if leaf.Left != nil {
		print(leaf.Left)
	}
	return
}
func main() {
	myTree := New()
	myTree.Add(9)
	myTree.Add(5)
	myTree.Add(13)
	fmt.Println(fmt.Sprintf("root %d, left %d, right %d", myTree.Root.Value, myTree.Root.Left, myTree.Root.Right))
	myTree.Print()
	fmt.Println("search", myTree.Search(5))
	myTree.Add(4)
	myTree.Add(8)
	myTree.Print()
	fmt.Println("search", myTree.Search(5))

	myTree.Delete(4)
	myTree.Print()

}
