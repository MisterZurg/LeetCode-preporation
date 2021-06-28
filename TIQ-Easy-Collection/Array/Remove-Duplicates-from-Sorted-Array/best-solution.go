package Remove_Duplicates_from_Sorted_Array

func removeDuplicatesBest(nums []int) int {
	j := 0
	x := 0

	for i := 0; i < len(nums); i++ {
		if i == 0 {
			nums[j] = nums[i]
			x++
		}
		if nums[i] != nums[j] {
			j++
			nums[j] = nums[i]
			x++
		}
	}
	return x
}
