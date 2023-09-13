package main

import (
	"fmt"
	"strings"
)

func minWindow(s string, t string) string {
	// Freq in t
	mp := [128]int{}
	have, want := 0, 0
	for _, c := range t {
		if mp[c] == 0 {
			want++
		}
		mp[c]++
	}

	cnt := [128]int{}

	beg, end := 0, 0

	d := strings.Repeat("@", len(s))

	// If current symbol in freq mp increment cnt freq
	// - if freq in mp == cnt increment have
	// - if have == want write update subst if it && try to minimize size of sbstr

	for end < len(s) {
		// fmt.Println(s[beg:end])
		if mp[s[end]] > 0 {
			cnt[s[end]]++
		}

		if mp[s[end]] == cnt[s[end]] && mp[s[end]] > 0 {
			have++
		}

		// Found substr where ABBC ABC
		for have == want {
			if len(d) >= end-beg+1 && have == want {
				d = s[beg : end+1] // <---
			}

			if mp[s[beg]] > 0 {
				cnt[s[beg]]--
				if mp[s[beg]] > cnt[s[beg]] {
					have--
				}
			}
			beg++
		}
		end++
	}

	if d == strings.Repeat("@", len(s)) {
		return ""
	}
	return d
}

func main() {
	fmt.Println(minWindow("ADOBECODEBANC", "ABC"))
	fmt.Println(minWindow("a", "a"))
	fmt.Println(minWindow("a", "aa"))
	fmt.Println(minWindow("aa", "aa"))
	// Output: "BANC
}

// https://leetcode.com/problems/minimum-window-substring/solutions/746521/a-visualized-first-principles-approach-for-better-understanding/
