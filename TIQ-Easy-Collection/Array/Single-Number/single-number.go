package Single_Number

func singleNumber(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	check := make(map[int]int)
	var ans int
	for _, number := range nums {
		check[number]++
	}

	for key, value := range check {
		if value == 1 {
			ans = key
			break
		}
	}
	return ans
}
