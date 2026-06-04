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

func balanceFactor(n *Node) int {
	return h(n.right) - h(n.left)
}

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
	bf := balanceFactor(n)
	if bf == 2 {
		if balanceFactor(n.right) < 0 {
			n.right = rotateRight(n.right)
		}
		return rotateLeft(n)
	}
	if bf == -2 {
		if balanceFactor(n.left) > 0 {
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
	} else {
		return root
	}
	return balance(root)
}

func inorder(root *Node) {
	if root == nil {
		return
	}
	inorder(root.left)
	fmt.Printf("key=%d h=%d bf=%d\n", root.key, root.height, balanceFactor(root))
	inorder(root.right)
}

func main() {
	var root *Node
	vals := []int{10, 20, 30, 40, 50, 25}
	fmt.Println("Вставка:", vals)
	for _, v := range vals {
		root = insert(root, v)
	}
	fmt.Println("AVL-дерево (симметричный обход):")
	inorder(root)
}
