impl Solution {
    pub fn is_anagram(s: String, t: String) -> bool {
        if s.len() != t.len() {
            return false;
        }

        // two arrays
        let mut cnt: [i32; 26] = [0; 26];

        for ch in 0..s.len() {
            cnt[s.as_bytes()[ch] as usize - 97] += 1;
            cnt[t.as_bytes()[ch] as usize - 97] -= 1;
        }

        cnt.iter().all(|c| *c == 0)
    }
}