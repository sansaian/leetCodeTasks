# Intuition
The problem of determining if a binary tree has a path from root to leaf that sums up to a given target suggests a depth-first search approach. By traversing each path and accumulating the sum, we can check if any path's sum equals the target sum when we reach the leaf nodes.

# Approach
The solution involves a recursive depth-first search (DFS) function `dfs` that traverses the tree and checks if the sum of the values along the path from the root to any leaf equals the target sum.
- We start from the root and keep adding the value of each node to a running sum (`curr`).
- At each node, we check if it's a leaf node (both left and right children are `nil`). If it is, we check if the running sum plus the current node's value equals the target sum.
- The process is repeated recursively for both the left and right subtrees.
- If any path meets the criteria, we return `true`; otherwise, `false`.

### Steps
1. Initialize a global variable `target` with the target sum.
2. Start the DFS from the root, with an initial current sum of 0.
3. In the DFS function:
    - Return `false` if the current node is `nil` (base case for leaf's child).
    - Check if the current node is a leaf (both children are `nil`) and if the sum up to this node equals the target.
    - Recursively call DFS for the left and right children, adding the current node's value to the running sum.
    - Return `true` if either the left or right subtree returns `true`.

# Complexity
- Time complexity: O(n), where n is the number of nodes in the tree. We potentially visit each node once in the DFS process.
- Space complexity: O(h), where h is the height of the tree. This is due to the recursive call stack used in DFS. In the worst case (a skewed tree), the space complexity can be O(n).

# Code
```go
var target = 0

func hasPathSum(root *TreeNode, targetSum int) bool {
    target = targetSum
    return dfs(root, 0)
}

func dfs(root *TreeNode, curr int) bool {
    if root == nil {
        return false
    }
    if root.Left == nil && root.Right == nil {
        return (root.Val + curr) == target
    }
    curr += root.Val
    left := dfs(root.Left, curr)
    right := dfs(root.Right, curr)
    return left || right
}
