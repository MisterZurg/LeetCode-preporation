package Plus_One

func plusOne(digits []int) []int {
	if len(digits)-1 == 0 && digits[len(digits)-1] == 9 {
		digits[len(digits)-1] += 1
		return []int{1, 0}
	} else if len(digits)-1 == 0 {
		digits[len(digits)-1] += 1
		return digits
	}

	for lastIndex := len(digits) - 1; lastIndex > 0; lastIndex-- {
		// [1,2,9] -> [1, 3, 0]
		if digits[lastIndex] == 9 {
			digits[lastIndex] = 0
			if digits[lastIndex-1] != 9 {
				digits[lastIndex] += 1
				break
			}
		} else {
			// [1,2,3] -> [1, 3, 4]
			digits[lastIndex] += 1
			break
		}
	}
	return digits
}
