package main

import "fmt"

type Node struct {
	key   int
	left  *Node
	right *Node
}

func newNode(key int) *Node { return &Node{key: key} }

func rotateRight(x *Node) *Node {
	y := x.left
	x.left = y.right
	y.right = x
	return y
}

func rotateLeft(x *Node) *Node {
	y := x.right
	x.right = y.left
	y.left = x
	return y
}

func splay(root *Node, key int) *Node {
	if root == nil || root.key == key {
		return root
	}
	if key < root.key {
		if root.left == nil {
			return root
		}
		if key < root.left.key {
			root.left.left = splay(root.left.left, key)
			root = rotateRight(root)
		} else if key > root.left.key {
			root.left.right = splay(root.left.right, key)
			if root.left.right != nil {
				root.left = rotateLeft(root.left)
			}
		}
		if root.left == nil {
			return root
		}
		return rotateRight(root)
	} else {
		if root.right == nil {
			return root
		}
		if key > root.right.key {
			root.right.right = splay(root.right.right, key)
			root = rotateLeft(root)
		} else if key < root.right.key {
			root.right.left = splay(root.right.left, key)
			if root.right.left != nil {
				root.right = rotateRight(root.right)
			}
		}
		if root.right == nil {
			return root
		}
		return rotateLeft(root)
	}
}

func insert(root *Node, key int) *Node {
	if root == nil {
		return newNode(key)
	}
	root = splay(root, key)
	if root.key == key {
		return root
	}
	n := newNode(key)
	if key < root.key {
		n.right = root
		n.left = root.left
		root.left = nil
	} else {
		n.left = root
		n.right = root.right
		root.right = nil
	}
	return n
}

func findMin(root *Node) *Node {
	for root.left != nil {
		root = root.left
	}
	return root
}

func delete(root *Node, key int) *Node {
	if root == nil {
		return nil
	}
	root = splay(root, key)
	if root.key != key {
		return root
	}
	if root.left == nil {
		return root.right
	}
	right := root.right
	root = splay(root.left, key)
	root.right = right
	return root
}

func inorder(root *Node) {
	if root == nil {
		return
	}
	inorder(root.left)
	fmt.Printf("%d ", root.key)
	inorder(root.right)
}

func main() {
	var root *Node
	for _, v := range []int{10, 20, 5, 15, 3, 25} {
		root = insert(root, v)
	}
	fmt.Print("После вставок: ")
	inorder(root)
	fmt.Println()
	fmt.Printf("Корень после вставок: %d\n", root.key)

	root = splay(root, 15)
	fmt.Printf("После splay(15), корень: %d\n", root.key)
	fmt.Print("Дерево: ")
	inorder(root)
	fmt.Println()

	root = delete(root, 15)
	fmt.Print("После удаления 15: ")
	inorder(root)
	fmt.Println()

	root = delete(root, 10)
	fmt.Print("После удаления 10: ")
	inorder(root)
	fmt.Println()
}
