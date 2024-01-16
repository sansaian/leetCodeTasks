package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func deepestLeavesSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	queue := []*TreeNode{root}
	var ans int
	for len(queue) > 0 {
		nodesInCurrentLevel := len(queue)
		curLevel := 0
		for i := 0; i < nodesInCurrentLevel; i++ {
			node := queue[0]
			//delete first element
			queue = queue[1:]
			curLevel += node.Val
			// Put the next level onto the queue
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		ans = curLevel
	}
	return ans
}
