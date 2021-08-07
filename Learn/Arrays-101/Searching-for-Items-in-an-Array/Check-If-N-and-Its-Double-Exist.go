package Searching_for_Items_in_an_Array

func checkIfExist(arr []int) bool {
	seen := make(map[int]bool)
	for _, num := range arr {
		if _, ok := seen[num*2]; ok {
			return true
		}
		if num%2 == 0 {
			if _, ok := seen[num/2]; ok {
				return true
			}
		}
		seen[num] = true
	}
	return false
}
