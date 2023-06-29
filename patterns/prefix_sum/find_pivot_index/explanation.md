# Intuition
The problem asks to find the pivot index of an array where the sum of all the numbers to the left of the index is equal to the sum of all the numbers to the right of the index. If no such index exists, we should return `-1`. This can be solved by first computing the total sum of the array, then iterating through the array and updating the left sum, and checking if the left sum equals the total sum minus the left sum and the current number.

# Approach
We initialize two variables, `sum` and `leftSum`, to `0`. `sum` is used to store the total sum of the array, and `leftSum` is used to store the sum of the numbers to the left of the current index.

We first calculate `sum` by iterating over the array and adding each value to `sum`.

Then, we iterate over the array again and for each index `i`, we check if `leftSum` equals `sum - leftSum - nums[i]`. If this condition is true, it means the sum of the numbers to the left of index `i` equals the sum of the numbers to the right of index `i`, and we return `i`.

If we iterate over the entire array and do not find any pivot index, we return `-1`.

# Complexity
- Time complexity: The time complexity for this algorithm is O(n), where `n` is the length of the `nums` array. This is because we perform two passes over the `nums` array.

- Space complexity: The space complexity is O(1), as we use a constant amount of space to store the `sum` and `leftSum` variables.

# Code
```go
func pivotIndex(nums []int) int {
    var sum,leftSum int
    for _,value:=range nums{
        sum+=value
    }
    for i:=range nums{
        if leftSum==sum-leftSum-nums[i]{
            return i
        }
        leftSum+=nums[i]
    }
    return -1
}
```