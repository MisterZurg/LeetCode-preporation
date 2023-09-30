package main

import "fmt"

type WordDictionary struct {
	root *WDTrieNode
}

func WordDictionaryConstructor() WordDictionary {
	return WordDictionary{
		root: new(WDTrieNode),
	}
}

func (this *WordDictionary) AddWord(word string) {
	curr := this.root
	for _, lt := range word {
		if curr.childrien[lt] == nil {
			curr.childrien[lt] = &WDTrieNode{}
		}
		curr = curr.childrien[lt]
	}
	curr.end = true
}

func (this *WordDictionary) Search(word string) bool {
	var dfs func(j int, node *WDTrieNode) bool
	dfs = func(j int, node *WDTrieNode) bool {
		if node == nil {
			return false
		}

		curr := node
		for i := j; i < len(word); i++ {
			lt := word[i]
			if lt == '.' {
				for _, child := range curr.childrien {
					if dfs(i+1, child) {
						return true
					}
				}
				return false
			} else {
				if curr.childrien[lt] == nil {
					return false
				}
				curr = curr.childrien[lt]
			}
		}
		return curr.end
	}
	return dfs(0, this.root)
}

type WDTrieNode struct {
	childrien [128]*WDTrieNode
	end       bool
}

func main() {
	wd := WordDictionaryConstructor()
	wd.AddWord("bad")
	wd.AddWord("dad")
	wd.AddWord("mad")
	fmt.Println(wd.Search("pad"))
	fmt.Println(wd.Search("bad"))
	fmt.Println(wd.Search(".ad"))
	fmt.Println(wd.Search("b.."))
}
