package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	queue := []*TreeNode{root}
	var ans [][]int
	var level int
	for len(queue) > 0 {
		level++
		nodesInCurrentLevel := len(queue)
		curlevel := make([]int, nodesInCurrentLevel)
		var node *TreeNode
		for i := 0; i < nodesInCurrentLevel; i++ {
			node = queue[0]
			queue = queue[1:]
			if (level % 2) == 0 {
				curlevel[nodesInCurrentLevel-1-i] = node.Val
			} else {
				curlevel[i] = node.Val
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		ans = append(ans, curlevel)
	}
	return ans
}
