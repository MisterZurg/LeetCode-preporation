package main

import "fmt"

func sortedSquares(nums []int) []int {
	N := len(nums)
	// Move positive pointer after zero
	positivePtr := 0
	for positivePtr < N && nums[positivePtr] < 0 {
		positivePtr++
	}

	negtivePtr := positivePtr - 1

	answerSortedSquares := make([]int, len(nums))
	// Used for iterate throug array
	counter := 0

	for negtivePtr >= 0 && N > positivePtr {
		negSqr := nums[negtivePtr] * nums[negtivePtr]
		posSqr := nums[positivePtr] * nums[positivePtr]

		if negSqr < posSqr {
			answerSortedSquares[counter] = negSqr
			negtivePtr--
		} else {
			answerSortedSquares[counter] = posSqr
			positivePtr++
		}
		counter++
	}

	// Corner cases
	for negtivePtr >= 0 {
		negSqr := nums[negtivePtr] * nums[negtivePtr]
		answerSortedSquares[counter] = negSqr
		negtivePtr--

		counter++
	}

	for positivePtr < N {
		posSqr := nums[positivePtr] * nums[positivePtr]
		answerSortedSquares[counter] = posSqr
		positivePtr++

		counter++
	}

	return answerSortedSquares
}

func main() {
	fmt.Println(sortedSquares([]int{-4, -1, 0, 3, 10}))
}
