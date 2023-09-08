package main

import "sort"

func threeSum(nums []int) [][]int {
	var badtrip [][]int

	if len(nums) < 3 {
		return badtrip
	}

	sort.Ints(nums)

	for k := 0; k < len(nums)-2; k++ {
		// skip repeated fixed letters
		if k > 0 && nums[k] == nums[k-1] {
			continue
		}

		i := k + 1
		j := len(nums) - 1

		for i < j {
			threesome := nums[k] + nums[i] + nums[j]
			if threesome == 0 {
				badtrip = append(badtrip, []int{nums[k], nums[i], nums[j]})
				j--

				// skip repeaded from end
				for i < j && nums[j] == nums[j+1] {
					j--
				}
			} else if threesome > 0 {
				j--
			} else {
				i++
			}
		}
	}

	return badtrip
}
