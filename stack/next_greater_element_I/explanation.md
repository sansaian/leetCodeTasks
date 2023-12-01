# Intuition
Initially, the problem of finding the next greater element for each element in an array suggests the use of a data structure that can efficiently track the elements and their next greater counterparts. A stack is a good fit for this purpose, as it can help maintain the order of elements and find the next greater element in an efficient manner.

# Approach
The solution involves using a stack and a hashmap. As we iterate through `nums2`, elements are pushed onto the stack until a greater element is found. When a greater element is encountered, it is recorded in the hashmap as the "next greater" for the elements in the stack that are smaller than it. For `nums1`, we then simply look up the hashmap to determine the next greater element for each of its values.

### Steps
1. Initialize a hashmap `mapstore` to store the mapping between elements from `nums1` and their next greater elements.
2. Iterate through `nums2`, using a stack to track elements.
    - For each element `v` in `nums2`, check the stack; if the top element of the stack is less than `v`, pop it from the stack and add to the hashmap `mapstore`, indicating that `v` is the next greater element.
    - Push `v` onto the stack.
3. Iterate through `nums1` and update each of its elements based on the values in `mapstore`.
    - If an element `v` from `nums1` has a corresponding element in `mapstore`, update it to the value from `mapstore`.
    - If not, set it to `-1` to denote no next greater element exists.

# Complexity
- Time complexity: O(n + m), where n is the size of `nums2`, and m is the size of `nums1`. Each array is traversed only once.
- Space complexity: O(n), where n is the size of `nums2`. The hashmap and stack can contain as many elements as there are in `nums2` in the worst case.

# Code
```go
func nextGreaterElement(nums1 []int, nums2 []int) []int {
    mapstore := make(map[int]int, len(nums1))
    var stack []int
    for _, v := range nums2 {
        for len(stack) > 0 && stack[len(stack)-1] < v {
            mapstore[stack[len(stack)-1]] = v
            stack = stack[:len(stack)-1]
        }
        stack = append(stack, v)
    }
    for i, v := range nums1 {
        if mapv, ok := mapstore[v]; ok {
            nums1[i] = mapv
            continue
        }
        nums1[i] = -1
    }
    return nums1
}
```