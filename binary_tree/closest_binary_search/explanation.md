# Problem Statement
Given the root of a binary search tree (BST) and a target value, the task is to find the value in the BST that is closest to the target. In case of a tie, return the smaller value.

## Example 1
- **Input**: root = [4,2,5,1,3], target = 3.714286
- **Output**: 4

## Example 2
- **Input**: root = [1], target = 4.428571
- **Output**: 1

# Approach
The solution iteratively traverses the binary search tree to find the value closest to the target. At each step, it decides whether to move left or right based on how the current node's value compares to the target. The `result` variable is updated whenever a closer value is found. If a value in the BST is equally close to the target as the current `result`, the smaller one among them is chosen as the new `result`.

## Steps
1. Initialize `result` with the value of the root node.
2. Traverse the tree starting from the root:
    - Update `result` if the current node's value is closer to the target than the current `result`. In case of a tie, prefer the smaller value.
    - Decide the next direction (left or right subtree) based on whether the target is greater or less than the current node's value.
3. Continue the traversal until the current node is `nil`.
4. Return `result` as the closest value to the target.

# Complexity Analysis
- **Time Complexity**: O(h), where h is the height of the BST. In the worst case, the algorithm might have to traverse from the root to the leaf, making the time complexity proportional to the height of the tree.
- **Space Complexity**: O(1), as the algorithm uses a fixed amount of space regardless of the input tree size.

# Code
```go
import "math"

// Definition for a binary tree node.
type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func closestValue(root *TreeNode, target float64) int {
    result := root.Val
    for root != nil {
        if math.Abs(float64(result) - target) > math.Abs(float64(root.Val) - target) ||
           (math.Abs(float64(result) - target) == math.Abs(float64(root.Val) - target) && result > root.Val) {
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
