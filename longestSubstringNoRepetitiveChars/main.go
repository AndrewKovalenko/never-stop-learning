package main

import "fmt"

func maxOf(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func searchSubstring(s string, foundSubstringLength int) int {
	if len(s) == 0 {
		return foundSubstringLength
	}

	substrLength := 0
	charactersFound := make(map[rune]bool)

	for i, char := range s {
		if charactersFound[char] {
			bestResultSoFar := maxOf(substrLength, foundSubstringLength)
			substrLengthInTheRestOfTheString := searchSubstring(s[i:], bestResultSoFar)

			return maxOf(substrLengthInTheRestOfTheString, bestResultSoFar)
		} else {
			charactersFound[char] = true
			substrLength = substrLength + 1
		}
	}

	return maxOf(substrLength, foundSubstringLength)
}

func lengthOfLongestSubstring(s string) int {
	return searchSubstring(s, 0)
}

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
	fmt.Println(lengthOfLongestSubstring("bbb"))
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
	fmt.Println(lengthOfLongestSubstring(""))
	fmt.Println(lengthOfLongestSubstring("aab"))
}
