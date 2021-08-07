package Rotate_Array

func rotate(nums []int, k int) {
	for i := 0; i < k; i++ {
		// Сохраняем последний элемент
		lastNum := nums[len(nums)-1]
		// Сдвигаем элементы на один вперед
		for j := len(nums) - 1; j > 0; j-- {
			nums[j] = nums[j-1]
		}
		nums[0] = lastNum
	}
}
