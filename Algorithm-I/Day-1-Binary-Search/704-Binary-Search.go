package main

import "fmt"

func search(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	// Nums already is sorted in ascending order
	// So we could use it!
	// Get idx of middle element
	for left <= right {
		m := (left + right) / 2

		if nums[m] == target {
			return m
		}

		// Move right border
		if nums[m] < target {
			left = m + 1
		} else {
			right = m - 1
		}

	}
	// If we didn't find any
	return -1
}

/*
Example 1:
Input: nums = [-1,0,3,5,9,12], target = 9
Output: 4
Explanation: 9 exists in nums and its index is 4

Example 2:
Input: nums = [-1,0,3,5,9,12], target = 2
Output: -1
Explanation: 2 does not exist in nums so return -1
*/
func main() {
	type TestCase struct {
		Nums   []int
		target int
	}
	var addTests = []TestCase{
		TestCase{
			Nums:   []int{-1, 0, 3, 5, 9, 12},
			target: 9,
		},
		TestCase{
			Nums:   []int{-1, 0, 3, 5, 9, 12},
			target: 2,
		},
	}

	for _, tc := range addTests {
		fmt.Println(search(tc.Nums, tc.target))
	}
}
