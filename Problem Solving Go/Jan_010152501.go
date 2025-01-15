func minimizeXor(num1 int, num2 int) int {
    num2Bits := bits.OnesCount(uint(num2)) // Count set bits in num2
	x := 0

	// Set the most significant bits of num1 to x
	for i := 31; i >= 0 && num2Bits > 0; i-- {
		if num1&(1<<i) != 0 {
			x |= (1 << i) // Set bit in x
			num2Bits--
		}
	}

	// If more bits are needed, set the least significant bits
	for i := 0; i < 32 && num2Bits > 0; i++ {
		if x&(1<<i) == 0 {
			x |= (1 << i) // Set bit in x
			num2Bits--
		}
	}

	return x
}