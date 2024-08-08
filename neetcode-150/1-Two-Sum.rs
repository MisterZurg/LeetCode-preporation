impl Solution {
    pub fn two_sum(nums: Vec<i32>, target: i32) -> Vec<i32> {
        use std::collections::HashMap;

        // num -> idx
        let mut hm: HashMap<i32, i32> = HashMap::new();

        // You may assume that each input
        // would have exactly one solution
        for (i, v) in nums.iter().enumerate() {
            match hm.get(&(target - *v)) {
                Some(&idx2) => return vec![i as i32, idx2],
                None => hm.insert(*v, i as i32),
            };
        }

        unreachable!();
    }
}