package main

import (
	"fmt"
	"math"
)

func lengthOfLongestSubstring(s string) int {
	maxLengthFound := 0
	start := 0
	charsFound := make(map[rune]bool)

	for i, char := range s {
		for charsFound[char] {
			startRune := rune(s[start])
			charsFound[startRune] = false
			start = start + 1
		}

		charsFound[char] = true
		length := i - start + 1
		maxLengthFound = int(math.Max(float64(length), float64(maxLengthFound)))
	}

	return int(math.Max(float64(maxLengthFound), 0))
}

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
	fmt.Println(lengthOfLongestSubstring("bbb"))
	fmt.Println(lengthOfLongestSubstring("pwpker"))
	fmt.Println(lengthOfLongestSubstring(""))
	fmt.Println(lengthOfLongestSubstring("aab"))
}
