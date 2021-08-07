package Contains_Duplicate

import "sort"

func containsDuplicate(nums []int) bool {
	sort.Ints(nums)
	for i := 1; i < len(nums); i++ {
		if nums[i-1] == nums[i] {
			return true
		}
	}
	// Прошлись по nums и все уникально
	return false
}
