package Rotate_Array

func rotateNotDumb(nums []int, k int) {
	k = k % len(nums)
	reverse(&nums, 0, len(nums)-1)
	reverse(&nums, 0, k-1)
	reverse(&nums, k, len(nums)-1)
}

func reverse(s *[]int, start, end int) {
	for start < end {
		tmp := (*s)[start]
		(*s)[start] = (*s)[end]
		(*s)[end] = tmp
		// Or we can do it:
		// (*s)[start],(*s)[end]=(*s)[end],(*s)[start]
		start++
		end--
	}
}
