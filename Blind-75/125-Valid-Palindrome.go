package main

import "strings"

//https://leetcode.com/problems/valid-palindrome/description/
func isPalindrome(s string) bool {
	n := len(s)
	// Two Pointers
	s = strings.ToLower(s)

	i := 0     // first symbol
	j := n - 1 // last symbol

	for i <= j {
		// skip non-alphanumeric characters from front
		for i <= j && !('a' <= s[i] && s[i] <= 'z') && !('0' <= s[i] && s[i] <= '9') {
			i++
		}
		// skip non-alphanumeric characters from back
		for i <= j && !('a' <= s[j] && s[j] <= 'z') && !('0' <= s[j] && s[j] <= '9') {
			j--
		}

		if i <= j && s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}

func main() {

}
