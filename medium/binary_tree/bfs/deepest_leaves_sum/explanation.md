# Intuition
The task is to find the sum of the values of the deepest leaves in a binary tree. This problem suggests a level-order traversal approach, where we process the tree level by level and keep track of the sum of values at each level. The sum of the last level processed, which will be the deepest, is the required answer.

# Approach
The solution involves using a queue to perform a level-order traversal (breadth-first search) of the binary tree. For each level:
- We count the number of nodes (`nodesInCurrentLevel`) to process at this level.
- Initialize a variable (`curLevel`) to store the sum of values of nodes at the current level.
- Process each node at this level, adding its value to `curLevel` and enqueueing its children for the next level.
- After processing all nodes at a level, update `ans` with `curLevel`. When the loop finishes, `ans` will hold the sum of values of the deepest leaves.

### Steps
1. If the root is `nil`, return 0 as there's no tree to process.
2. Initialize a queue with the root node and a variable `ans` to hold the sum.
3. While the queue is not empty:
    - Determine the number of nodes at the current level (`nodesInCurrentLevel`).
    - Initialize `curLevel` to 0 to calculate the sum of values at this level.
    - Iterate through nodes at this level, updating `curLevel` with the node's value and enqueueing its children.
    - After processing all nodes at this level, update ans with curLevel.
4. Return ans, which now contains the sum of the deepest leaves.

### Complexity
- Time complexity: O(n), where n is the number of nodes in the tree. Each node is visited exactly once during the level-order traversal.
- Space complexity: O(m), where m is the maximum number of nodes at any level in the tree. This space is required for the queue used in the level-order traversal, particularly when storing nodes of the widest level.

# Code
```
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func deepestLeavesSum(root *TreeNode) int {
	if root == nil { return 0} 
	queue := []*TreeNode{root} 
    var ans int
	for len(queue) > 0 { 
		nodesInCurrentLevel := len(queue) 
        curLevel:=0
		for i := 0; i < nodesInCurrentLevel; i++ { 
			node := queue[0] 
			//delete first element
			queue = queue[1:] 
            curLevel+=node.Val
			// Put the next level onto the queue 
			if node.Left != nil { 
				queue = append(queue, node.Left) 
			} 
			if node.Right != nil { 
				queue = append(queue, node.Right) 
			} 
		}
        ans=curLevel
	}
    return ans
}
```