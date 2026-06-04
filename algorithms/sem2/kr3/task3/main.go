package main

import "fmt"

type Node struct {
	key    int
	height int
	left   *Node
	right  *Node
}

func h(n *Node) int {
	if n == nil {
		return 0
	}
	return n.height
}

func updateHeight(n *Node) {
	l, r := h(n.left), h(n.right)
	if l > r {
		n.height = l + 1
	} else {
		n.height = r + 1
	}
}

func bf(n *Node) int { return h(n.right) - h(n.left) }

func rotateRight(y *Node) *Node {
	x := y.left
	y.left = x.right
	x.right = y
	updateHeight(y)
	updateHeight(x)
	return x
}

func rotateLeft(x *Node) *Node {
	y := x.right
	x.right = y.left
	y.left = x
	updateHeight(x)
	updateHeight(y)
	return y
}

func balance(n *Node) *Node {
	updateHeight(n)
	if bf(n) == 2 {
		if bf(n.right) < 0 {
			n.right = rotateRight(n.right)
		}
		return rotateLeft(n)
	}
	if bf(n) == -2 {
		if bf(n.left) > 0 {
			n.left = rotateLeft(n.left)
		}
		return rotateRight(n)
	}
	return n
}

func insert(root *Node, key int) *Node {
	if root == nil {
		return &Node{key: key, height: 1}
	}
	if key < root.key {
		root.left = insert(root.left, key)
	} else if key > root.key {
		root.right = insert(root.right, key)
	}
	return balance(root)
}

func findMin(root *Node) *Node {
	for root.left != nil {
		root = root.left
	}
	return root
}

func removeMin(root *Node) *Node {
	if root.left == nil {
		return root.right
	}
	root.left = removeMin(root.left)
	return balance(root)
}

func delete(root *Node, key int) *Node {
	if root == nil {
		return nil
	}
	if key < root.key {
		root.left = delete(root.left, key)
	} else if key > root.key {
		root.right = delete(root.right, key)
	} else {
		l, r := root.left, root.right
		if r == nil {
			return l
		}
		min := findMin(r)
		min.right = removeMin(r)
		min.left = l
		return balance(min)
	}
	return balance(root)
}

func inorder(root *Node) {
	if root == nil {
		return
	}
	inorder(root.left)
	fmt.Printf("key=%d bf=%d\n", root.key, bf(root))
	inorder(root.right)
}

func main() {
	var root *Node
	for _, v := range []int{10, 20, 30, 40, 50, 25} {
		root = insert(root, v)
	}
	fmt.Println("Исходное AVL-дерево:")
	inorder(root)

	root = delete(root, 20)
	fmt.Println("\nПосле удаления 20:")
	inorder(root)

	root = delete(root, 10)
	fmt.Println("\nПосле удаления 10:")
	inorder(root)
}
