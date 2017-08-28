package main

import (
	"fmt"
)

func lengthOfLongestSubstring(s string) int {
	runes := []rune(s)
	runeMap := make(map[rune]int)
	longest := 0
	preLength := 0
	for i, rn := range runes {
		var length int
		if val, ok := runeMap[rn]; !ok || val < i-preLength {
			length = preLength + 1
		} else {
			length = i - val
		}
		if length > longest {
			longest = length
		}
		preLength = length
		runeMap[rn] = i
	}
	return longest
}

func main() {
	fmt.Printf("%#v\n", lengthOfLongestSubstring("abcabcbb"))
	fmt.Printf("%#v\n", lengthOfLongestSubstring("bbbbb"))
	fmt.Printf("%#v\n", lengthOfLongestSubstring("pwwkew"))
}
