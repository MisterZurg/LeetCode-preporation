package main

import (
	"math"
	"unicode"
)

func myAtoiBest(s string) int {
	index := 0
	// Read in and ignore any leading whitespace
	for index < len(s) && s[index] == ' ' {
		index++
	}
	// Terminate if there are no characters left
	if index == len(s) {
		return 0
	}
	// Check if the number starts with a + or -
	signum := 1
	if s[index] == '+' {
		index++ // just ignore it
	} else if s[index] == '-' {
		signum = -1
		index++
	}
	// Use a number larger than and int to prevent overflow
	var result int64
	// While there are digits to read, read them
	for index < len(s) && unicode.IsDigit(rune(s[index])) {
		result *= 10
		result += int64(s[index] - '0')
		index++
		if signum == -1 && -result < math.MinInt32 {
			return math.MinInt32
		} else if signum == 1 && result > math.MaxInt32 {
			return math.MaxInt32
		}
	}
	// clamp the number within -2^31 and 2^31 - 1
	result *= int64(signum)
	return int(result)
}
