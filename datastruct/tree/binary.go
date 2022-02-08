package main

import "fmt"

type BinaryTree struct {
	Root *Leaf
}

type Leaf struct {
	Left  *Leaf
	Right *Leaf
	Value int
}

func New() *BinaryTree {
	return &BinaryTree{
		Root: nil,
	}
}

func (b *BinaryTree) Delete(val int) {}

func (b *BinaryTree) Add(val int) {
	b.Root = add(b.Root, val)
}

func add(leaf *Leaf, val int) *Leaf {
	if leaf == nil {
		return &Leaf{
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

func (b *BinaryTree) Search(val int) *Leaf {
	return search(b.Root, val)
}

func search(leaf *Leaf, val int) *Leaf {
	if leaf.Value == val || leaf == nil {
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

func main() {
	myTree := New()
	myTree.Add(9)
	myTree.Add(5)
	myTree.Add(13)
	fmt.Println("search", myTree.Search(5))
	myTree.Add(4)
	myTree.Add(8)
	fmt.Println("search", myTree.Search(5))
	fmt.Println(myTree.Root)
}
