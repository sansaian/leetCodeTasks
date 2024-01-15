package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func largestValues(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var ans []int
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		nodesInCurrentLevel := len(queue)
		min := -9223372036854775808
		for i := 0; i < nodesInCurrentLevel; i++ {
			node := queue[0]
			queue = queue[1:]
			min = max(node.Val, min)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		ans = append(ans, min)
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
