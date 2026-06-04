package main

import "fmt"

type Node struct {
	val         int
	left, right *Node
}

func isMaxHeap(n *Node) bool {
	if n == nil {
		return true
	}
	if n.left != nil && n.val < n.left.val {
		return false
	}
	if n.right != nil && n.val < n.right.val {
		return false
	}
	return isMaxHeap(n.left) && isMaxHeap(n.right)
}

func isMinHeap(n *Node) bool {
	if n == nil {
		return true
	}
	if n.left != nil && n.val > n.left.val {
		return false
	}
	if n.right != nil && n.val > n.right.val {
		return false
	}
	return isMinHeap(n.left) && isMinHeap(n.right)
}

func isHeap(root *Node) bool {
	return isMaxHeap(root) || isMinHeap(root)
}

func buildTree(arr []int, i int) *Node {
	if i >= len(arr) || arr[i] == -1 {
		return nil
	}
	return &Node{
		val:   arr[i],
		left:  buildTree(arr, 2*i+1),
		right: buildTree(arr, 2*i+2),
	}
}

func result(b bool) string {
	if b {
		return "куча"
	}
	return "не куча"
}

func main() {
	t1 := buildTree([]int{10, 5, 8, 2, 3}, 0)
	fmt.Printf("[10,5,8,2,3]: %s\n", result(isHeap(t1)))

	t2 := buildTree([]int{5, 10, 8, 2, 3}, 0)
	fmt.Printf("[5,10,8,2,3]: %s\n", result(isHeap(t2)))

	t3 := buildTree([]int{1, 3, 2, 7, 5}, 0)
	fmt.Printf("[1,3,2,7,5]: %s\n", result(isHeap(t3)))

	t4 := buildTree([]int{3, 1, 5, 2, 4}, 0)
	fmt.Printf("[3,1,5,2,4]: %s\n", result(isHeap(t4)))
}
