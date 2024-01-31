package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxAncestorDiff(root *TreeNode) int {

	return findMax(root, root.Val, root.Val)
}

func findMax(root *TreeNode, minVal, maxVal int) int {
	if root == nil {
		return maxVal - minVal
	}
	minVal = min(minVal, root.Val)
	maxVal = max(maxVal, root.Val)
	leftDiff := findMax(root.Left, minVal, maxVal)
	rightDiff := findMax(root.Right, minVal, maxVal)
	return max(leftDiff, rightDiff)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
