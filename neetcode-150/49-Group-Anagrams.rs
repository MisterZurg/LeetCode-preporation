impl Solution {
    pub fn group_anagrams(strs: Vec<String>) -> Vec<Vec<String>> {
        use std::collections::{HashSet, HashMap};

        // let mut groups = HashSet::new();

        // anogram -> anograms idx
        let mut mp: HashMap<[u8; 26], usize> = HashMap::new();
        let mut anagrams: Vec<Vec<String>> = Vec::new();
        for s in strs {
            let mut cnt: [u8; 26] = [0; 26];
            for b in s.bytes() {
                cnt[(b - b'a') as usize] += 1;
            }

            match mp.get(&cnt) {
                Some(idx) => anagrams[*idx].push(s),
                None => {
                    anagrams.push(vec![s]);
                    mp.insert(cnt, anagrams.len() - 1);
                }
            }
        }
        return anagrams;
    }
}