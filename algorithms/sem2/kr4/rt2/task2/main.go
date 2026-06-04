package main

import "fmt"

type TrieNode struct {
	children [26]*TrieNode
	isEnd    bool
}

type Trie struct {
	root *TrieNode
}

func NewTrie() *Trie { return &Trie{root: &TrieNode{}} }

func (t *Trie) Insert(word string) {
	node := t.root
	for _, ch := range word {
		idx := ch - 'a'
		if node.children[idx] == nil {
			node.children[idx] = &TrieNode{}
		}
		node = node.children[idx]
	}
	node.isEnd = true
}

func (t *Trie) countPrefixes(word string) int {
	node := t.root
	count := 0
	for i, ch := range word {
		idx := ch - 'a'
		if node.children[idx] == nil {
			break
		}
		node = node.children[idx]
		if node.isEnd && i < len(word)-1 {
			count++
		}
	}
	return count
}

func main() {
	words := []string{"a", "ab", "abc", "abcd", "abcdef", "bcd"}

	trie := NewTrie()
	for _, w := range words {
		trie.Insert(w)
	}

	bestWord := ""
	bestCount := -1

	for _, w := range words {
		cnt := trie.countPrefixes(w)
		fmt.Printf("  %q: префиксов = %d\n", w, cnt)
		if cnt > bestCount {
			bestCount = cnt
			bestWord = w
		}
	}

	fmt.Printf("\nОтвет: %q (префиксов: %d)\n", bestWord, bestCount)
}
