package Introduction

import "fmt"

func findNumbers(nums []int) int {
	count := 0
	// iterate through each and convert to string
	// if string has even length then increase counter
	for _, n := range nums {
		if len(fmt.Sprintf("%d", n))%2 == 0 {
			count++
		}
	}
	return count
}

/*
func findNumbers(nums []int) int {
    evens := 0
    for _, number := range nums {
        if howManyDigits(number) % 2 == 0 {
            evens++
        }
    }
    return evens
}

func howManyDigits(number int) int {
    // 10^5
    if number < 10 {
        return 1
    }
    if number < 100 {
        return 2
    }
    if number < 1000 {
        return 3
    }
    if number < 10000 {
        return 4
    }
    if number < 100000 {
        return 5
    }
    return 6
}

*/
