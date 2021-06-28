package First_Unique_Character_in_a_String

import "strings"

func firstUniqChar(s string) int {
	for ind, letter := range s {
		if strings.Count(s, string(letter)) == 1 {
			return ind
		}
	}

	return -1
}
