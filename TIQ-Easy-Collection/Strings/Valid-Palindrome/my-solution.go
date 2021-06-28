package main

import (
	"fmt"
	"strings"
)
func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	var probablyPalindrome = make([]rune, 0)
	for _, letter := range s {
		if '0' <= rune(letter) && rune(letter) <= '9' ||  'a' <= rune(letter) && rune(letter) <= 'z'{
			probablyPalindrome = append(probablyPalindrome, letter)
		}
	}
	answ := string(probablyPalindrome)
	for i:=0 ; i < len(answ) / 2; i++ {
		if answ[i] != answ[len(answ) - 1 - i] {
			return false
		}
	}
	return true
}

func main()  {
	fmt.Println(isPalindrome("A man, a plan, a canal: Panama")) // Output: true
}