package main

import (
	"strconv"
	"unicode/utf8"
)

func compress(chars []byte) int {
	storeIdx := 0
	i := 0
	for i < len(chars) {
		j := i + 1
		for j < len(chars) && chars[i] == chars[j] {
			j++
		}
		chars[storeIdx] = chars[i]
		storeIdx++
		if count := j - i; count > 1 {
			// Calculate the number of digits
			digits := utf8.RuneCountInString(strconv.Itoa(count))
			storeIdx += digits
			// Start storing the digits in reverse
			for n, k := count, storeIdx-1; n != 0; n, k = n/10, k-1 {
				chars[k] = byte(n%10 + '0')
			}
		}
		i = j
	}
	return storeIdx
}
