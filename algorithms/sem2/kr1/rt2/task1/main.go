package main

import "fmt"

type Node struct {
	data int
	next *Node
}

func checkList(head *Node) bool {
	slow, fast := head, head
	for fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next
		if slow == fast {
			return false
		}
	}

	expected := head
	cur := head
	for cur != nil {
		if cur != expected {
			return false
		}
		expected = expected.next
		cur = cur.next
	}
	return true
}

func buildList(vals []int) *Node {
	if len(vals) == 0 {
		return nil
	}
	head := &Node{data: vals[0]}
	cur := head
	for _, v := range vals[1:] {
		cur.next = &Node{data: v}
		cur = cur.next
	}
	return head
}

func printList(head *Node) {
	for cur := head; cur != nil; cur = cur.next {
		fmt.Printf("%d -> ", cur.data)
	}
	fmt.Println("nil")
}

func main() {
	nodes := make([]*Node, 5)
	for i, v := range []int{1, 2, 3, 4, 5} {
		nodes[i] = &Node{data: v}
	}
	for i := 0; i < 4; i++ {
		nodes[i].next = nodes[i+1]
	}
	fmt.Print("Корректный список: ")
	printList(nodes[0])
	fmt.Println("Проверка:", checkList(nodes[0]))

	nodes2 := make([]*Node, 5)
	for i, v := range []int{1, 2, 3, 4, 5} {
		nodes2[i] = &Node{data: v}
	}
	nodes2[0].next = nodes2[1]
	nodes2[1].next = nodes2[2]
	nodes2[2].next = nodes2[4]
	nodes2[3].next = nil
	fmt.Println("\nСписок с проскоком (1->2->3->5, пропущен 4):")
	fmt.Print("Обход: ")
	for cur := nodes2[0]; cur != nil; cur = cur.next {
		fmt.Printf("%d -> ", cur.data)
	}
	fmt.Println("nil")
	fmt.Println("Проверка:", checkList(nodes2[0]))

	nodes3 := make([]*Node, 5)
	for i, v := range []int{1, 2, 3, 4, 5} {
		nodes3[i] = &Node{data: v}
	}
	for i := 0; i < 4; i++ {
		nodes3[i].next = nodes3[i+1]
	}
	nodes3[4].next = nodes3[2]
	fmt.Println("\nСписок с перескоком (цикл 3->4->5->3):")
	fmt.Println("Проверка:", checkList(nodes3[0]))
}
