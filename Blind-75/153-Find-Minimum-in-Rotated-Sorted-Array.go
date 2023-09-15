package main

import "math"

func findMin(nums []int) int {
	l := 0
	r := len(nums)
	minimum := math.MaxInt
	for l < r {
		m := (l + r) / 2
		// 4 1 2 3
		if nums[m] >= nums[l] {
			// In rotated part
			minimum = min(minimum, nums[l])
			l = m + 1
		} else {
			minimum = min(minimum, nums[m])
			r = m
		}
	}
	return minimum
}

//func min(a, b int) int {
//	if a < b {
//		return a
//	}
//	return b
//}
