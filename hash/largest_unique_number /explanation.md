# Intuition
Given a list of integers, we want to identify the largest integer that occurs only once in the list. We can leverage hashmaps to efficiently count occurrences of numbers and subsequently identify the largest unique number.

# Approach
1. **Initialize the Hashmap**: A hashmap (or a dictionary in other languages) called `store` is initialized to keep track of the occurrence of each number in the `nums` list.

2. **Count Occurrences**: We loop through the `nums` list and increment the corresponding count in the `store` for every number encountered.

3. **Identify Largest Unique Number**: We initialize a variable `max` to `-1` (which will be our return value if no unique number is found). We then loop through the `store`, and for every number that occurs only once (i.e., its count is 1), we update the `max` if the current number is greater than the current value of `max`.

# Complexity
- **Time Complexity**: The time complexity is O(n), where n is the number of elements in the `nums` list. This is because we do a single iteration over the list to fill our hashmap, and another iteration over the keys in the hashmap (which is at most n).

- **Space Complexity**: The space complexity is O(n), where n is the number of elements in the `nums` list. This accounts for the storage used by our hashmap.

# Code
```go
func largestUniqueNumber(nums []int) int {
    // Hashmap to store occurrence counts
    store := make(map[int]int, len(nums))
    
    // Counting occurrences
    for _, v := range nums {
        store[v] = store[v] + 1
    } 
    
    max := -1  // Default value if no unique number is found
    // Identifying the largest unique number
    for k, v := range store {
        if v != 1 {
            continue
        }
        if max < k {
            max = k
        }
    }
    return max
}
