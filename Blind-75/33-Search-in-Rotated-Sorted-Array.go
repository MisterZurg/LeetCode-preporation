package main

func search(nums []int, target int) int {
	l := 0
	r := len(nums) - 1
	// [4,5,6,[7],0,1,2]

	for l <= r {
		m := (l + r) / 2

		if nums[m] == target {
			return m
		}

		// If our pivot > target
		// - [4,5,[6],7,0,1,2]
		//            |- target

		// fmt.Println(l, m, r)

		if nums[l] <= nums[m] {
			// We are in the LEFT part
			if nums[m] < target || target < nums[l] {
				l = m + 1
			} else {
				r = m - 1
			}
		} else { // We are in the RIGHT part
			// - [4,5,6,7,0,[1],2]
			//            |- targetg
			if target < nums[m] || target > nums[r] {
				r = m - 1
			} else {
				l = m + 1
			}
		}
	}

	return -1
}
