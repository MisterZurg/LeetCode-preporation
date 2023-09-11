package main

func characterReplacement(s string, k int) int {
	// Let's do some more NeetCode today

	// Sliding window
	mp := [26]int{}
	beg := 0 // in
	end := 0
	counter := 0 // frequncy of most repeated letters in a string
	d := 0

	for end < len(s) {
		// Get the frequency of the most repeated letter in a string
		idx := s[end] - 'A'
		mp[idx]++
		for _, freq := range mp {
			counter = max(counter, freq)
		}

		// Check that the window size - freqMostRepeadLetter <= k
		windowSize := end - beg + 1
		if windowSize-counter <= k {
			d = max(d, end-beg+1)
		} else {
			idx = s[beg] - 'A'
			mp[idx]--
			beg++
		}

		end++
	}
	return d
}

//func max(a, b int) int {
//	if a > b {
//		return a
//	}
//	return b
//}
