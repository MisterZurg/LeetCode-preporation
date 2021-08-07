package Introduction

func findMaxConsecutiveOnes(nums []int) int {
	current, max := 0, 0
	for _, num := range nums {
		if num != 0 {
			current++
		} else {
			if current > max {
				max = current
			}
			current = 0
		}
	}
	if current > max {
		max = current
	}
	return max
}
