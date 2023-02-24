package main

import tree "github.com/leetCodeTasks/datastruct/binary_tree"

func main() {
	mytree := sortedArrayToBST([]int{-10, -3, 0, 5, 9})
	mytree.Print()
}

func sortedArrayToBST(nums []int) *tree.TreeNode {
	return ToBST(nums, 0, len(nums)-1)
}

func ToBST(nums []int, start, finish int) *tree.TreeNode {
	if start > finish {
		return nil
	}
	mid := start + (finish-start)/2
	root := &tree.TreeNode{Val: nums[mid]}

	root.Left = ToBST(nums, start, mid-1)
	root.Right = ToBST(nums, mid+1, finish)

	return root
}
