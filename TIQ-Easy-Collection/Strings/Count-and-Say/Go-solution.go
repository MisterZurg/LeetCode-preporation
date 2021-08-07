package main

import (
	"fmt"
	"strconv"
)

func countAndSay(n int) string {
	// Default case
	if n == 1 {
		return "1"
	} else if n == 2 {
		return "11"
	}
	// Else two pointers strategy
	prevAns := countAndSay(n - 1)
	ans := ""
	pCounter := 1
	for idx := 1; idx < len(prevAns); idx++ {
		if prevAns[idx-1] == prevAns[idx] {
			pCounter++
		} else {
			ans += strconv.FormatInt(int64(pCounter), 10)
			ans += string(prevAns[idx-1])
			pCounter = 1
		}
		// fmt.Printf("pCounter: %d | idx : %d\n", pCounter, idx)
	}

	ans += strconv.FormatInt(int64(pCounter), 10)

	ans += string(prevAns[len(prevAns)-1])
	return ans
}

func main() {
	//fmt.Println(countAndSay(1))  //     1
	//fmt.Println(countAndSay(2))  //     11
	fmt.Println(countAndSay(3)) //     21
	fmt.Println(countAndSay(4)) //    1211
	//fmt.Println(countAndSay(5))  //     111221
	//fmt.Println(countAndSay(6))  //     312211
	//fmt.Println(countAndSay(7))  //    13112221
	//fmt.Println(countAndSay(8))  //    1113213211
	//fmt.Println(countAndSay(9))  //    31131211131221
	//fmt.Println(countAndSay(10)) //     13211311123113112211)
}
