package main

import "math"

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func closestValue(root *TreeNode, target float64) int {
	result := root.Val
	for root != nil {
		if math.Abs(float64(result)-target) > math.Abs(float64(root.Val)-target) ||
			math.Abs(float64(result)-target) == math.Abs(float64(root.Val)-target) &&
				result > root.Val {
			result = root.Val
		}

		if target > float64(root.Val) {
			root = root.Right
		} else {
			root = root.Left
		}
	}
	return result
}
