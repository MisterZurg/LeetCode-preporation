package Deleting_Items_From_an_Array

func removeDuplicates(nums []int) int {
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
