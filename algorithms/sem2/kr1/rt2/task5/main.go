package main

import "fmt"

type Node struct {
	digit int
	next  *Node
}

func fromSlice(digits []int) *Node {
	if len(digits) == 0 {
		return nil
	}
	head := &Node{digit: digits[0]}
	cur := head
	for _, d := range digits[1:] {
		cur.next = &Node{digit: d}
		cur = cur.next
	}
	return head
}

func printList(head *Node) {
	for cur := head; cur != nil; cur = cur.next {
		fmt.Printf("%d ", cur.digit)
	}
	fmt.Println()
}

func addLists(l1, l2 *Node) *Node {
	dummy := &Node{}
	cur := dummy
	carry := 0
	for l1 != nil || l2 != nil || carry != 0 {
		sum := carry
		if l1 != nil {
			sum += l1.digit
			l1 = l1.next
		}
		if l2 != nil {
			sum += l2.digit
			l2 = l2.next
		}
		carry = sum / 10
		cur.next = &Node{digit: sum % 10}
		cur = cur.next
	}
	return dummy.next
}

func main() {
	l1 := fromSlice([]int{2, 4, 3})
	l2 := fromSlice([]int{5, 6, 4})

	fmt.Print("L1: ")
	printList(l1)
	fmt.Print("L2: ")
	printList(l2)

	result := addLists(l1, l2)
	fmt.Print("Сумма: ")
	printList(result)

	l3 := fromSlice([]int{9, 9, 9})
	l4 := fromSlice([]int{1})
	fmt.Print("\nL3: ")
	printList(l3)
	fmt.Print("L4: ")
	printList(l4)
	fmt.Print("Сумма: ")
	printList(addLists(l3, l4))
}
