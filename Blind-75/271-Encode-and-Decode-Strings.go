package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

// https://leetcode.com/problems/encode-and-decode-strings/
// https://www.lintcode.com/problem/659/
func encode(words []string) string {
	// string concatanation entirely
	var sb strings.Builder
	for _, word := range words {
		sb.WriteString(fmt.Sprintf("%d#%s", utf8.RuneCountInString(word), word))
	}
	return sb.String()
}

func decode(str string) []string {
	var decoded []string
	var sb strings.Builder

	for i := 0; i < utf8.RuneCountInString(str); i++ {
		if unicode.IsDigit(rune(str[i])) {
			sb.WriteRune(rune(str[i]))
		} else if str[i] == '#' {
			wSize, _ := strconv.Atoi(sb.String())
			decoded = append(decoded, str[i+1:i+1+wSize])
			i += wSize
			sb = strings.Builder{}
		}
	}
	return decoded
}

func main() {
	input := [][]string{
		{"lint", "code", "love", "you"},
		{"we", "say", ":", "yes"},
		{"#", "##", "####", "####"},
	}

	for _, str := range input {
		fmt.Println(str)
		fmt.Println(decode(encode(str)))
	}
}
