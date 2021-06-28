package Valid_Anagram

import "strings"

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	for _, letter := range s {
		if strings.Count(s, string(letter)) != strings.Count(t, string(letter)) {
			return false
		}
	}
	return true
}
