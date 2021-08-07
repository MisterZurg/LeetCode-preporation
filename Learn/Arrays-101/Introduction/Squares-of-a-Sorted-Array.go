package Introduction

func duplicateZeros(arr []int) {
	zeros := 0
	// Count zeros
	for _, number := range arr {
		if number == 0 {
			zeros++
		}
	}
	// [1 0 2]
	for i := len(arr) - 1; i >= 0; i-- {
		if arr[i] == 0 {
			if i+zeros < len(arr) {
				arr[i+zeros] = 0
			}
			zeros--
		}
		if i+zeros < len(arr) {
			arr[i+zeros] = arr[i]
		}
	}
}
