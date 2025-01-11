package main

import (
	"fmt"
)

func canConstruct(s string, k int) bool {
	// If k is greater than the length of the string, it's impossible to form k palindromes
	if k > len(s) {
		return false
	}

	// Map to count the frequency of each character
	charCount := make(map[rune]int)
	for _, c := range s {
		charCount[c]++
	}

	oddCount := 0

	// Count the number of characters with odd frequencies
	for _, count := range charCount {
		if count%2 != 0 {
			oddCount++
		}
	}

	// To form k palindromes, there can be at most k characters with odd frequencies
	return oddCount <= k
}