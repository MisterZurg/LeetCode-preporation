package Searching_for_Items_in_an_Array

func validMountainArray(arr []int) bool {
	// Default cases
	if len(arr) == 1 || len(arr) == 2 {
		return false
	}
	reachedPeak := false
	for idx := 1; idx < len(arr)-1; idx++ {
		// Перегиб [ 2 9 5 ]
		if !reachedPeak && arr[idx-1] < arr[idx] && arr[idx] > arr[idx+1] {
			reachedPeak = true
		}
		if (arr[idx-1] == arr[idx]) || (arr[idx] == arr[idx+1]) {
			return false
		}

		if (!reachedPeak && arr[idx-1] > arr[idx]) || (reachedPeak && arr[idx] < arr[idx+1]) {
			return false
		}
	}
	return reachedPeak
}
