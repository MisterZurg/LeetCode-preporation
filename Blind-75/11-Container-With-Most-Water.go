package main

func maxArea(height []int) int {
	i := 0
	j := len(height) - 1

	best := 0

	for i < j {
		dist := j - i
		minWall := min(height[i], height[j])

		best = max(best, minWall*dist)

		if i < j && height[i] < height[j] {
			i++
		} else {
			j--
		}
	}
	return best
}

// LeetCode Go 1.18 -_-
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
