# Function: lengthOfLongestSubstring
The `lengthOfLongestSubstring` function calculates the length of the longest substring without repeating characters within a given string `s`.

## Parameters
- `s` (string): The input string in which we are looking to find the length of the longest substring without repeating characters.

## Return Value
- (int): Returns the length of the longest non-repeating character substring found in `s`.

## Methodology
This function implements the sliding window technique to dynamically resize the window of characters under consideration without repeating characters. We use two pointers, `left` and `right`, which represent the beginning and end of our current window, and a hashmap `store` to keep track of the counts of each character within the window.

### Steps
1. Initialize a hashmap `store` that maps a byte (`uint8`) to an integer, representing characters in the substring and their frequency.
2. Initialize two integer pointers, `left` and `right`, to traverse the string. `left` represents the start of the current substring, and `right` represents the end.
3. Initialize an integer `result` to keep track of the length of the longest substring found so far.
4. Iterate over the string with the `right` pointer, extending the window to the right one character at a time.
5. When a character is repeated (`store[r] > 1`), shrink the window from the left by increasing the `left` pointer and updating the frequency of the characters passed.
6. Update `result` to hold the maximum length found between the current result and the new window size (`right-left+1`).
7. Continue this process until the `right` pointer has traversed the entire string.
8. Return `result` which holds the length of the longest substring without repeating characters.

## Auxiliary Function: max
A helper function `max` is defined to return the maximum of two integers, used to update `result` with the maximum length found.

## Complexity Analysis
- **Time Complexity**: O(n) where `n` is the length of the string `s`. Each character is processed once when the `right` pointer moves and at most once again when the `left` pointer moves.
- **Space Complexity**: O(min(m, n)) where `n` is the length of the string `s` and `m` is the size of the character set (if `s` consists of all possible characters, then `m` would be the character set size). This space is used by the hashmap that stores the characters in the current window.

## Code
```go
func lengthOfLongestSubstring(s string) int {
    store := make(map[uint8]int)
    var left, right, result int

    for right < len(s) {
        r := s[right]
        store[r]++
        for store[r] > 1 {
            l := s[left]
            store[l]--
            left++
        }
        result = max(result, right-left+1)
        right++
    }
    return result
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
