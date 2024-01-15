# Intuition
The task is to find the largest value in each level of a binary tree. This can be approached by a level-order traversal, where we keep track of the maximum value at each level of the tree.

# Approach
The solution involves using a queue for a level-order traversal (breadth-first search) of the tree. For each level, we:
- Initialize a variable (`min`) with a very small value to track the largest value at the current level.
- Process each node at the current level, updating `min` to the maximum value found.
- Enqueue the left and right children of each node to the queue for the next level's processing.
- After processing all nodes at a level, append min to the answer array (ans) as it represents the largest value for that level.
- Continue this process until the queue is empty, indicating that all levels of the tree have been processed.
# Steps
1. Check if the root is nil. If so, return an empty slice as there's no tree to process.
2. Initialize an empty slice ans to store the largest values and a queue for level-order traversal, starting with the root node.
3. While the queue is not empty:
    - Determine the number of nodes at the current level (nodesInCurrentLevel).
    - Initialize min with a very small value.
    - Iterate through the nodes at the current level:
        - Update min with the maximum value between min and the current node's value.
        - Enqueue the left and right children of the current node (if they exist).
    - After processing all nodes at this level, append min to ans.
4. Return the ans slice containing the largest values from each level.
# Complexity
- Time complexity: O(n), where n is the number of nodes in the tree. Each node is processed exactly once.
- Space complexity: Space complexity: O(m), where m is the maximum number of nodes at any level in the tree. This space is required for the queue used in the level-order traversal.

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
func largestValues(root *TreeNode) []int {
    if root==nil{
        return nil
    }
    var ans []int
    var nodesInCurrentLevel int
    queue := []*TreeNode{root} 
    	for len(queue) > 0 { 
		nodesInCurrentLevel = len(queue)
        min:=-922337203685
		for i := 0; i < nodesInCurrentLevel; i++ { 
			node := queue[0] 
		    queue = queue[1:] 
            min=max(node.Val,min)
			if node.Left != nil { 
				queue = append(queue, node.Left) 
			} 
			if node.Right != nil { 
			    queue = append(queue, node.Right) 
			}
		} 
        ans=append(ans,min)
	} 
    return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
```