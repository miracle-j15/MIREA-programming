package main

import "fmt"

type Node struct {
	key    int
	left   *Node
	right  *Node
	parent *Node
	balance int
}

func insert(root *Node, key int) *Node {
	if root == nil {
		return &Node{key: key}
	}
	if key < root.key {
		root.left = insert(root.left, key)
		root.left.parent = root
	} else if key > root.key {
		root.right = insert(root.right, key)
		root.right.parent = root
	}
	return root
}

func height(root *Node) int {
	if root == nil {
		return 0
	}
	l, r := height(root.left), height(root.right)
	if l > r {
		return l + 1
	}
	return r + 1
}

func calcBalance(root *Node) {
	if root == nil {
		return
	}
	root.balance = height(root.right) - height(root.left)
	calcBalance(root.left)
	calcBalance(root.right)
}

func printInorder(root *Node) {
	if root == nil {
		return
	}
	printInorder(root.left)
	fmt.Printf("key=%d balance=%d\n", root.key, root.balance)
	printInorder(root.right)
}

func main() {
	var root *Node
	for _, v := range []int{10, 5, 15, 3, 7, 12, 20} {
		root = insert(root, v)
	}
	calcBalance(root)
	fmt.Println("Узлы с показателями баланса (симметричный обход):")
	printInorder(root)
}
