package main

func searchInsert(nums []int, target int) int {
	left := 0
	right := len(nums) - 1

	for left <= right {
		m := (left + right) / 2

		if nums[m] == target {
			return m
		}

		if nums[m] < target {
			left = m + 1
		} else {
			right = m - 1
		}
	}
	// Answer if nothing was found
	return left
}
