# Intuition
The challenge is to perform a level-order traversal of a binary tree but with a twist: the nodes at each level should be arranged in a zigzag manner. This suggests using a queue for level-order traversal while alternating the order in which values are added to the output list at each level.

# Approach
The solution involves using a queue for level-order traversal and a variable `level` to keep track of the current level. For each level:
- Count the number of nodes (`nodesInCurrentLevel`) to process.
- Create a slice `curlevel` to store node values at this level.
- Iterate through nodes at this level, adjusting the insertion index based on whether the level is odd or even (for zigzag pattern).
- Enqueue the left and right children of each node.
- After processing all nodes at this level, append `curlevel` to the result (`ans`).

### Steps
1. Check if the root is `nil`. If so, return an empty slice as there's no tree to process.
2. Initialize a queue with the root node, an empty slice `ans` for the result, and a variable `level` starting at 0.
3. While the queue is not empty:
    - Increment `level`.
    - Determine the number of nodes at the current level (`nodesInCurrentLevel`).
    - Create a slice `curlevel` with a length equal to `nodesInCurrentLevel`.
    - For each node in the level, determine its position in `curlevel` based on the zigzag pattern.
    - Enqueue the left and right children of each node.
    - Append `curlevel` to `ans`.
4. Return `ans`, which now contains the zigzag level order traversal of the tree.

# Complexity
- Time complexity: O(n), where n is the number of nodes in the tree. Each node is visited exactly once.
- Space complexity: O(m), where m is the maximum number of nodes at any level in the tree. This space is required for the queue and the `curlevel` slice at each level.

# Code
```go
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
