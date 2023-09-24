package main

import "fmt"

type Trie struct {
	root *TrieNode
}

// TrieConstructor is renamed Constructor for Trie creation
func TrieConstructor() Trie {
	return Trie{
		root: &TrieNode{},
	}
}

func (this *Trie) Insert(word string) {
	curr := this.root
	for _, letter := range word {
		// fmt.Println("inserting:", letter)
		if curr.children[letter-65] == nil {
			curr.children[letter-65] = &TrieNode{
				val: letter,
			}
		}
		curr = curr.children[letter-65]
	}
	curr.end = true
}

func (this *Trie) Search(word string) bool {
	curr := this.root
	for _, letter := range word {
		if curr.children[letter-65] == nil {
			return false
		}
		curr = curr.children[letter-65]
	}
	return curr.end
}

func (this *Trie) StartsWith(prefix string) bool {
	curr := this.root
	for _, letter := range prefix {
		if curr.children[letter-65] == nil {
			return false
		}
		curr = curr.children[letter-65]
	}
	return true
}

type TrieNode struct {
	// TrieNode is an alphabet letter so we can replace
	val      rune // output helper
	children [128]*TrieNode
	end      bool
}

//func NewTrieNode(val uint) *TrieNode {
//
//}

func main() {
	obj := TrieConstructor()
	obj.Insert("LeetCode")
	param_2 := obj.Search("LeetCode")
	param_3 := obj.StartsWith("LeetCod")
	fmt.Println(param_2, param_3)
	//fmt.Println(obj.rootNodes)
	// fmt.Println('A', 'a')
	fmt.Println(obj.root)
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
