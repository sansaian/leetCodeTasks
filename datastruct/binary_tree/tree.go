package binary_tree

import "fmt"

type TreeNode struct {
	Left  *TreeNode
	Right *TreeNode
	Value int
}

func New(rootValue int) *TreeNode {
	return &TreeNode{
		Value: rootValue,
	}
}

func (b *TreeNode) Delete(val int) {
	delete(b, nil, val)
}

func delete(node *TreeNode, upNode *TreeNode, val int) {
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

func (b *TreeNode) Add(val int) {
	add(b, val)
}

func add(leaf *TreeNode, val int) *TreeNode {
	if leaf == nil {
		return &TreeNode{
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

func (b *TreeNode) Search(val int) *TreeNode {
	return search(b, val)
}

func search(leaf *TreeNode, val int) *TreeNode {
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

func (b *TreeNode) Print() {
	print(b)
}

func print(leaf *TreeNode) {
	fmt.Println(leaf.Value)
	if leaf.Right != nil {
		print(leaf.Right)
	}
	if leaf.Left != nil {
		print(leaf.Left)
	}
	return
}
