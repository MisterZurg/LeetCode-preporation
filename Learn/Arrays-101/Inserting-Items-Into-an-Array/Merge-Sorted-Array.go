package Inserting_Items_Into_an_Array

func merge(nums1 []int, m int, nums2 []int, n int) {
	// Поскольку мы знаем что nums1 длинны n + m
	// то все сводится к применению стратегии двух поинтеров
	// Поскольку второй массив меньше или равен длинны первого массива будем идти по нему
	// Индекс в первом массиве
	// Нужно сравнить элементы в двух массивах
	// [1 2 3 0 0 0]
	//  2 3 4
	/* Через стандартную библиотеку Го
	   for i:=m ; i< len(nums1);i++ {
	       nums1[i] = nums2[i - m]
	   }
	   sort.Ints(nums1)
	*/

	// Алгоритмически
	last := m + n - 1
	for n > 0 && m > 0 {
		if nums1[m-1] < nums2[n-1] {
			nums1[last] = nums2[n-1]
			n--
		} else {
			nums1[last] = nums1[m-1]
			m--
		}
		last--
	}
	//
	for n > 0 {
		nums1[last] = nums2[n-1]
		n--
		last--
	}
}
