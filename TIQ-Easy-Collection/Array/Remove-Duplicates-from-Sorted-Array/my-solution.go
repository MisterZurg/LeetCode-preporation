package Remove_Duplicates_from_Sorted_Array

func removeDuplicates(nums []int) int {
	var i = 0
	for ind := 1; ind < len(nums); ind++ {
		if nums[ind-1] != nums[ind] {
			i += 1
			nums[i] = nums[ind]
		}
	}
	return i + 1
}
