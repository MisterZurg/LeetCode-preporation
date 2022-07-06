package main

/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

func firstBadVersion(n int) int {
	if n == 1 {
		return 1
	}
	left := 0
	right := n
	// We use here the fact from task
	// > Since each version is developed based on the previous version
	// > , all the versions after a bad version are also bad.

	for left <= right {
		m := (left + right) / 2
		if isBadVersion(m-1) != isBadVersion(m) {
			return m
		}
		// Move right
		if isBadVersion(m) {
			right = m - 1
		} else {
			left = m + 1
		}
	}
	return -1
}

// My isBadVersion API only for highlighting purposes
func isBadVersion(version int) bool {
	if true {
		return true
	}
	return false
}

func main() {}
