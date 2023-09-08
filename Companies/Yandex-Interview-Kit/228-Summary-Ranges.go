package main

import (
	"fmt"
	"strconv"
)

func summaryRanges(nums []int) []string {
	if len(nums) == 0 {
		return []string{}
	} else if len(nums) == 1 {
		return []string{
			strconv.Itoa(nums[0]),
		}
	}

	result := make([]string, 0)
	curr := nums[0]

	for i := range nums {
		if i == len(nums)-1 || nums[i]+1 != nums[i+1] {
			if nums[i] != curr {
				tmp := strconv.Itoa(curr) + "->" + strconv.Itoa(nums[i])
				result = append(result, tmp)
			} else {
				result = append(result, strconv.Itoa(curr))
			}
			if i != len(nums)-1 {
				curr = nums[i+1]
			}
		}
	}

	return result
}

func main() {
	fmt.Println(summaryRanges([]int{-1}))
}
