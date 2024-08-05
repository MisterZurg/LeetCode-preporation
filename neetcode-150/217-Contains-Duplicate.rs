use std::collections::HashSet;

impl Solution {
    pub fn contains_duplicate(nums: Vec<i32>) -> bool {
        let mut hs : HashSet<i32> = HashSet::new();

        for n in nums {
            if hs.contains(&n) {
                return true;
            }

            hs.insert(n);
        }

        false
    }
}