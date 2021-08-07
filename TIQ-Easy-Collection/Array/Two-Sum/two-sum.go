package Two_Sum

func twoSum(nums []int, target int) []int {
	// Создадим мапу[число]индекс
	parsedNumbers := make(map[int]int)
	// Проходимся по nums
	for idx, number := range nums {
		// Проверяем наличие в мапе числа,
		//сумма которого с текущим даст нам target
		if _, ok := parsedNumbers[target-number]; ok {
			return []int{parsedNumbers[target-number], idx}
		} else {
			parsedNumbers[number] = idx
		}
	}
	return nil
}
