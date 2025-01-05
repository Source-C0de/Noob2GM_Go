package main

func shiftCharacter (s string, shifts  [][]int) string{
	n := len(s)
	shiftAmount := make([]int, n+1)

	// Apply shift intervals to the shiftAmount array
	for _, shift := range shifts {
		start, end, direction := shift[0], shift[1], shift[2]
		if direction == 1 {
			shiftAmount[start]++
			shiftAmount[end+1]--
		} else {
			shiftAmount[start]--
			shiftAmount[end+1]++
		}
	}

	// Calculate cumulative shifts
	currentShift := 0
	for i := 0; i < n; i++ {
		currentShift += shiftAmount[i]
		shiftAmount[i] = currentShift
	}

	// Apply the shifts to the string
	result := []rune(s)
	for i := 0; i < n; i++ {
		shift := (int(result[i]-'a') + shiftAmount[i]) % 26
		if shift < 0 {
			shift += 26
		}
		result[i] = rune('a' + shift)
	}

	return string(result)
}