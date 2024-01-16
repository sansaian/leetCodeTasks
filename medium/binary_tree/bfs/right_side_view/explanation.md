# Intuition
The goal is to create a view of a binary tree from the right side, which means capturing the rightmost node at each level of the tree. This suggests using a level-order traversal (breadth-first search) approach, where we ensure that for each level, the last node encountered in the traversal is recorded.

# Approach
The `rightSideView` function uses a queue to perform a level-order traversal of the tree. The process involves:
- Enqueuing the root node initially.
- Iterating over each level of the tree. For each level:
    - Track the number of nodes (`nodesInCurrentLevel`).
    - Traverse through these nodes, updating a variable `prev` with the value of the current node.
    - After finishing each level, the value in `prev` will be the rightmost node of that level, which is then appended to the `ans` slice.
- The loop continues until all levels are processed and the queue is empty.

### Steps
1. Return `nil` if the root is `nil`, as there's no tree to view.
2. Initialize a slice `ans` to store the right side view.
3. Use a queue to perform a level-order traversal, starting with the root.
4. While the queue is not empty:
    - Determine the number of nodes at the current level (`nodesInCurrentLevel`).
    - Iterate through these nodes:
        - Dequeue each node and update `prev` with its value.
        - Enqueue the left and right children of the dequeued node (if they exist).
    - After completing the level, append `prev` to `ans`.
5. Return `ans`, which contains the right side view of the tree.

# Complexity
- Time complexity: O(n), where n is the number of nodes in the tree. Each node is processed exactly once.
- Space complexity: O
  (m), where m is the maximum number of nodes at any level of the tree. This space is required for the queue used in the level-order traversal.

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
func rightSideView(root *TreeNode) []int {
    if root==nil{
        return nil
    }
    var ans []int
    queue := []*TreeNode{root} 
    	for len(queue) > 0 { 
		nodesInCurrentLevel := len(queue)
        prev:=0
		for i := 0; i < nodesInCurrentLevel; i++ { 
			node := queue[0] 
		    queue = queue[1:] 
            prev=node.Val
			if node.Left != nil { 
				queue = append(queue, node.Left) 
			} 
			if node.Right != nil { 
			    queue = append(queue, node.Right) 
			}
		} 
        ans=append(ans,prev)
	} 
    return ans
}
```