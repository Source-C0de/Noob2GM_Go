func minimumLength(s string) int {
    charFrequencyMap := make(map[rune]int)
	for _, currentChar := range s {
		charFrequencyMap[currentChar]++
	}

	// Step 2: Calculate the number of characters to delete
	deleteCount := 0
	for _, frequency := range charFrequencyMap {
		if frequency%2 == 1 {
			// If frequency is odd, delete all except one
			deleteCount += frequency - 1
		} else {
			// If frequency is even, delete all except two
			deleteCount += frequency - 2
		}
	}

	// Step 3: Return the minimum length after deletions
	return len(s) - deleteCount
}