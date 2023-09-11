package main

import "fmt"

// Using Killer Template
func lengthOfLongestSubstring(s string) int {
	//  ASCII encodes 128 specified characters
	var mp [128]bool
	start, end := 0, 0
	// We don't have counter here
	d := 0

	// Drop init  part here

	for end < len(s) {
		// Remove repeated chars
		for mp[s[end]] != false {
			mp[s[start]] = false
			start++
		}
		mp[s[end]] = true
		/* counter condition */
		//increase begin to make it invalid/valid again

		/* update d here if finding maximum */
		d = max(d, end-start+1)
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

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb")) //== 3)
	fmt.Println(lengthOfLongestSubstring("bbbbb"))    // == 1)
	fmt.Println(lengthOfLongestSubstring("pwwkew"))   //== 3)
	fmt.Println(lengthOfLongestSubstring("bb"))       //== 1)
	fmt.Println(lengthOfLongestSubstring(" "))        //== 1)
}

/* // Using HashMap
func lengthOfLongestSubstring(s string) int {
	mp := make(map[byte]int) // letter -> idx
	begin := 0
	counter := 0

	for end := range s {
		// Check if repeated symbol still in window
		_, ok := mp[s[end]]
		for ok {
			delete(mp, s[begin])
			begin++
			// Go limitation
			_, ok = mp[s[end]]
		}
		mp[s[end]] = end
		counter = max(counter, end-begin+1)
	}
	return counter
}
*/
