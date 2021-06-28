package Reverse_Integer

func reverse(x int) int {
	var answer int = x % 10
	x /= 10
	for ; x != 0; {
		if x == 0 {
			break
		}
		answer = answer*10 + x%10
		x /= 10
	}
	// math.MinInt32 && math.MaxInt32
	if -2147483648 < answer && answer < 2147483647 {
		return answer
	}
	return 0
}
