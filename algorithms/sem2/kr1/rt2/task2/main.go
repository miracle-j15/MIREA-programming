package main

import "fmt"

type Node struct {
	data int
	next *Node
}

func partition(head *Node, x int) (*Node, *Node) {
	var lessHead, lessTail, greaterHead, greaterTail *Node

	for cur := head; cur != nil; cur = cur.next {
		if cur.data < x {
			if lessHead == nil {
				lessHead = cur
				lessTail = cur
			} else {
				lessTail.next = cur
				lessTail = cur
			}
		} else {
			if greaterHead == nil {
				greaterHead = cur
				greaterTail = cur
			} else {
				greaterTail.next = cur
				greaterTail = cur
			}
		}
	}

	if lessTail != nil {
		lessTail.next = nil
	}
	if greaterTail != nil {
		greaterTail.next = nil
	}

	return lessHead, greaterHead
}

func printList(head *Node) {
	for cur := head; cur != nil; cur = cur.next {
		fmt.Printf("%d -> ", cur.data)
	}
	fmt.Println("nil")
}

func main() {
	vals := []int{1, 4, 3, 2, 5, 2}
	head := &Node{data: vals[0]}
	cur := head
	for _, v := range vals[1:] {
		cur.next = &Node{data: v}
		cur = cur.next
	}

	var x int
	fmt.Scan(&x)

	l, r := partition(head, x)
	fmt.Print("Меньше: ")
	printList(l)
	fmt.Print("Больше/равно: ")
	printList(r)
}
