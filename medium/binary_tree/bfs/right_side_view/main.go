package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var ans []int
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		nodesInCurrentLevel := len(queue)
		prev := 0
		for i := 0; i < nodesInCurrentLevel; i++ {
			node := queue[0]
			queue = queue[1:]
			prev = node.Val
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		ans = append(ans, prev)
	}
	return ans
}
