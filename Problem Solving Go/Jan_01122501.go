

package main

import (
	"fmt"
)

func canBeValid(s string, locked string) bool {
	n := len(s)
	if n%2 != 0 {
		return false // A valid parentheses string must have an even length
	}

	// Left-to-right check: Ensure enough '(' to balance ')'
	open, free := 0, 0
	for i := 0; i < n; i++ {
		if locked[i] == '1' {
			if s[i] == '(' {
				open++
			} else {
				open--
			}
		} else {
			free++
		}

		if open+free < 0 {
			return false // Not enough '(' to balance ')'
		}
	}

	// Right-to-left check: Ensure enough ')' to balance '('
	close := 0
	free = 0
	for i := n - 1; i >= 0; i-- {
		if locked[i] == '1' {
			if s[i] == ')' {
				close++
			} else {
				close--
			}
		} else {
			free++
		}

		if close+free < 0 {
			return false // Not enough ')' to balance '('
		}
	}

	return true
}

func main() {
	s := "()))"
	locked := "0100"

	if canBeValid(s, locked) {
		fmt.Println("True")
	} else {
		fmt.Println("False")
	}
}
