package main

// O(n)
func rotate(nums []int, k int) {
	N := len(nums)
	rotated := make([]int, N)

	for i := range nums {
		idxForRotation := (i + k) % N
		rotated[idxForRotation] = nums[i]
	}

	for i := range nums {
		nums[i] = rotated[i]
	}
}

func main() {
	rotate([]int{1, 2, 3, 4, 5, 6, 7}, 3)
}
