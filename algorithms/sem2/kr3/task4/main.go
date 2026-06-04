package main

import "fmt"

const (
	Red   = "Red"
	Black = "Black"
)

type Node struct {
	key   int
	color string
	left  *Node
	right *Node
}

func insertBST(root *Node, key int) *Node {
	if root == nil {
		return &Node{key: key, color: Red}
	}
	if key < root.key {
		root.left = insertBST(root.left, key)
	} else if key > root.key {
		root.right = insertBST(root.right, key)
	}
	return root
}

func colorBlack(root *Node) string { return Black }

func blackHeight(root *Node) int {
	if root == nil {
		return 1
	}
	l := blackHeight(root.left)
	r := blackHeight(root.right)
	if l != r || l == -1 {
		return -1
	}
	add := 0
	if root.color == Black {
		add = 1
	}
	return l + add
}

func checkRB(root *Node) (bool, string) {
	if root == nil {
		return true, ""
	}
	if root.color != Black {
		return false, "корень не чёрный"
	}
	var check func(*Node) bool
	check = func(n *Node) bool {
		if n == nil {
			return true
		}
		if n.color == Red {
			lc := "Black"
			rc := "Black"
			if n.left != nil {
				lc = n.left.color
			}
			if n.right != nil {
				rc = n.right.color
			}
			if lc == Red || rc == Red {
				return false
			}
		}
		return check(n.left) && check(n.right)
	}
	if !check(root) {
		return false, "красный узел имеет красного потомка"
	}
	if blackHeight(root) == -1 {
		return false, "неодинаковая чёрная высота путей"
	}
	return true, ""
}

func inorder(root *Node) {
	if root == nil {
		return
	}
	inorder(root.left)
	fmt.Printf("key=%d color=%s\n", root.key, root.color)
	inorder(root.right)
}

func main() {
	var root *Node
	for _, v := range []int{10, 5, 15, 3, 7} {
		root = insertBST(root, v)
	}
	root.color = Black
	root.left.color = Black
	root.right.color = Black
	root.left.left.color = Red
	root.left.right.color = Red

	fmt.Println("Дерево:")
	inorder(root)
	ok, reason := checkRB(root)
	if ok {
		fmt.Println("Свойства КЧ-дерева выполнены")
	} else {
		fmt.Println("Нарушение:", reason)
	}

	root2 := insertBST(nil, 10)
	root2.color = Black
	root2.left = &Node{key: 5, color: Red}
	root2.left.left = &Node{key: 3, color: Red}
	fmt.Println("\nДерево с нарушением (красный → красный):")
	inorder(root2)
	ok2, reason2 := checkRB(root2)
	if ok2 {
		fmt.Println("Свойства выполнены")
	} else {
		fmt.Println("Нарушение:", reason2)
	}
}
