package main

import "fmt"

type Node struct {
	value int
	prev  *Node
	next  *Node
}

func intersect(h1, h2 *Node) *Node {
	dummy := &Node{}
	tail := dummy
	a, b := h1, h2
	for a != nil && b != nil {
		if a.value == b.value {
			n := &Node{value: a.value}
			if tail != dummy {
				tail.next = n
				n.prev = tail
			}
			tail = n
			a = a.next
			b = b.next
		} else if a.value < b.value {
			a = a.next
		} else {
			b = b.next
		}
	}
	if dummy.next != nil {
		dummy.next.prev = nil
	}
	return dummy.next
}

func build(vals []int) *Node {
	if len(vals) == 0 {
		return nil
	}
	head := &Node{value: vals[0]}
	cur := head
	for _, v := range vals[1:] {
		n := &Node{value: v, prev: cur}
		cur.next = n
		cur = n
	}
	return head
}

func printList(head *Node) {
	for cur := head; cur != nil; cur = cur.next {
		fmt.Printf("%d ", cur.value)
	}
	fmt.Println()
}

func main() {
	l1 := build([]int{2, 2, 3, 5, 7, 7, 9})
	l2 := build([]int{2, 2, 2, 3, 3, 7, 8})

	fmt.Print("L1: ")
	printList(l1)
	fmt.Print("L2: ")
	printList(l2)

	result := intersect(l1, l2)
	fmt.Print("Пересечение: ")
	printList(result)
}
