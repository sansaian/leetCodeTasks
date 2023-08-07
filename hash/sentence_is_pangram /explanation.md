# Intuition
The problem is asking us to determine if the input sentence is a pangram, i.e., it contains every letter of the alphabet at least once. To solve this problem, we can iterate through each character in the sentence and store each unique character in a data structure. If the total number of unique characters is 26 (the number of letters in the English alphabet), we know that the sentence is a pangram.

# Approach
The approach in this implementation is to use a map data structure, where each character in the sentence is a key. We iterate over each character in the sentence, adding new characters to the map. This ensures each character only appears once in the map, since map keys are unique. Finally, we check if the size of the map is 26. If so, this means that the sentence is a pangram; otherwise, it's not.

# Complexity
- Time complexity: The time complexity is O(n), where n is the length of the sentence. This is because we're iterating through each character in the sentence once.
- Space complexity: The space complexity is also O(n), where n is the length of the sentence. This is the space used by the map to store unique characters. In the worst case, if each character in the sentence is unique, we'd have to store every character in the map. Note that since the map's size won't exceed 26 (the number of letters in the alphabet), we could also argue that the space complexity is O(1), or constant.

# Code
```go
func checkIfPangram(sentence string) bool {
	store:=make(map[rune]struct{})
	for _,ch:=range sentence{
		if _,ok:=store[ch];!ok{
			store[ch]=struct{}{}
		}
	}
	return len(store)==26
}
