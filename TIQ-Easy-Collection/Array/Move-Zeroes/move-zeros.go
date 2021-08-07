package Move_Zeroes

func moveZeroes(nums []int) {
	if len(nums) == 1 {
		return
	}
	// nums = [0,1,0,3,12]
	// Ищем ноль
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == 0 {
			// Ищем число
			for j := i + 1; j < len(nums); j++ {
				if nums[j] != 0 {
					// Если находим то меняем их местами
					// и прерываем цикл по числам
					nums[j], nums[i] = nums[i], nums[j]
					break
				}
			}
		}
	}
}
