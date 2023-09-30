package main

import (
	"strconv"
)

func compress(chars []byte) int {
	i, j := 1, 0
	charsPtr := 0

	for i < len(chars) {
		if chars[i] == chars[i-1] {
			i++
		} else {
			chars[charsPtr] = chars[j]
			charsPtr++
			if i-j > 1 {
				number := strconv.Itoa(i - j)
				for _, dg := range number {
					chars[charsPtr] = byte(dg)
					charsPtr++
				}
			}
			j = i
			i++
		}
	}

	chars[charsPtr] = chars[j]
	charsPtr++
	if i-j > 1 {
		number := strconv.Itoa(i - j)
		for _, dg := range number {
			chars[charsPtr] = byte(dg)
			charsPtr++
		}
	}

	return charsPtr
}

// notInplace
/*
func compress(chars []byte) int {
	i, j := 1, 0
	var sb strings.Builder
	for i < len(chars) {
		// fmt.Println(chars[i])
		if chars[i] == chars[i-1] {
			i++
		} else {
			sb.WriteByte(chars[j])
			if i-j > 1 {
				number := strconv.Itoa(i - j)
				// fmt.Println(number)
				for _, dg := range number {
					sb.WriteRune(dg)
				}
				j = i
			}
			i++
		}
	}

	sb.WriteByte(chars[j])
	if i-j > 1 {
		number := strconv.Itoa(i - j)
		for dg := range number {
			sb.WriteByte(number[dg])
		}
	}
	fmt.Println(sb.String())
	return len(sb.String())
}

func main() {
	fmt.Println(compress([]byte{'a', 'a', 'b', 'b', 'c', 'c', 'c'}))
}
*/
