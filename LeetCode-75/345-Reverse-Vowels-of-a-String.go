package LeetCode_75

func reverseVowels(s string) string {
	// two pointers

	letters := []rune(s)

	VOWELS := make(map[rune]bool)

	VOWELS['a'] = true
	VOWELS['e'] = true
	VOWELS['i'] = true
	VOWELS['o'] = true
	VOWELS['u'] = true
	// nd they can appear in both lower and upper cases, more than once.
	VOWELS['A'] = true
	VOWELS['E'] = true
	VOWELS['I'] = true
	VOWELS['O'] = true
	VOWELS['U'] = true

	l, r := 0, len(s)-1
	for l < r {
		if _, ok := VOWELS[letters[l]]; !ok {
			l++
			continue
		}

		if _, ok := VOWELS[letters[r]]; !ok {
			r--
			continue
		}

		letters[l], letters[r] = letters[r], letters[l]
		l++
		r--
	}
	return string(letters)
}
