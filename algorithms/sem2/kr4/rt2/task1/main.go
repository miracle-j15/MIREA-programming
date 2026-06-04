package main

import "fmt"

type TrieNode struct {
	children [26]*TrieNode
	isEnd    bool
}

type Trie struct {
	root *TrieNode
}

func NewTrie() *Trie {
	return &Trie{root: &TrieNode{}}
}

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

func (t *Trie) Search(word string) bool {
	node := t.root
	for _, ch := range word {
		idx := ch - 'a'
		if node.children[idx] == nil {
			return false
		}
		node = node.children[idx]
	}
	return node.isEnd
}

func deleteHelper(node *TrieNode, word string, depth int) bool {
	if node == nil {
		return false
	}
	if depth == len(word) {
		node.isEnd = false
		return !node.hasChildren()
	}
	idx := word[depth] - 'a'
	if deleteHelper(node.children[idx], word, depth+1) {
		node.children[idx] = nil
		return !node.isEnd && !node.hasChildren()
	}
	return false
}

func (n *TrieNode) hasChildren() bool {
	for _, child := range n.children {
		if child != nil {
			return true
		}
	}
	return false
}

func (t *Trie) Delete(word string) {
	deleteHelper(t.root, word, 0)
}

func main() {
	trie := NewTrie()
	words := []string{"apple", "app", "application", "apply", "banana"}
	for _, w := range words {
		trie.Insert(w)
	}

	fmt.Println("До удаления:")
	for _, w := range words {
		fmt.Printf("  search(%q) = %v\n", w, trie.Search(w))
	}

	trie.Delete("app")
	fmt.Println("\nПосле удаления \"app\":")
	for _, w := range words {
		fmt.Printf("  search(%q) = %v\n", w, trie.Search(w))
	}

	trie.Delete("apple")
	fmt.Println("\nПосле удаления \"apple\":")
	fmt.Printf("  search(\"apple\")       = %v\n", trie.Search("apple"))
	fmt.Printf("  search(\"application\") = %v\n", trie.Search("application"))
	fmt.Printf("  search(\"apply\")       = %v\n", trie.Search("apply"))
}
