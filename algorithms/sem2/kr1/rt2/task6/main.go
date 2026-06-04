package main

import "fmt"

type Node struct {
	data int
	next *Node
}

func moveEvensToFront(head *Node) *Node {
	var evenHead, evenTail, oddHead, oddTail *Node
	for cur := head; cur != nil; cur = cur.next {
		if cur.data%2 == 0 {
			if evenHead == nil {
				evenHead = cur
				evenTail = cur
			} else {
				evenTail.next = cur
				evenTail = cur
			}
		} else {
			if oddHead == nil {
				oddHead = cur
				oddTail = cur
			} else {
				oddTail.next = cur
				oddTail = cur
			}
		}
	}
	if evenTail != nil {
		evenTail.next = oddHead
	}
	if oddTail != nil {
		oddTail.next = nil
	}
	if evenHead != nil {
		return evenHead
	}
	return oddHead
}

func printList(head *Node) {
	for cur := head; cur != nil; cur = cur.next {
		fmt.Printf("%d -> ", cur.data)
	}
	fmt.Println("nil")
}

func main() {
	vals := []int{5, 2, 8, 3, 1, 6}
	head := &Node{data: vals[0]}
	cur := head
	for _, v := range vals[1:] {
		cur.next = &Node{data: v}
		cur = cur.next
	}

	fmt.Print("До: ")
	printList(head)

	head = moveEvensToFront(head)
	fmt.Print("После: ")
	printList(head)
}
