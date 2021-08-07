package Intersection_of_Two_Arrays_II

func intersect(nums1 []int, nums2 []int) []int {
	// Я вижу решение через map[int]int ,
	// где ключ это собственно число в массиве, а значение сколько раз он встретился
	numbers := make(map[int]int)
	ans := make([]int, 0)
	if len(nums1) < len(nums2) {
		// Просто наполняем мапу числами из первого массива
		for _, num1 := range nums1 {
			numbers[num1]++
		}
		for _, num2 := range nums2 {
			if val, ok := numbers[num2]; ok && 0 < val {
				numbers[num2]--
				ans = append(ans, num2)
			}
		}

	} else {
		// Просто наполняем мапу числами из второго массива
		for _, num2 := range nums2 {
			numbers[num2]++
		}
		for _, num1 := range nums1 {
			if val, ok := numbers[num1]; ok && 0 < val {
				numbers[num1]--
				ans = append(ans, num1)
			}
		}
	}
	return ans
}
