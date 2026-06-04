package main

import "fmt"

type Node struct {
	value int
	next  *Node
	prev  *Node
}

type CircularList struct {
	head *Node
}

func (l *CircularList) Add(value int) {
	newNode := &Node{value: value}
	if l.head == nil {
		newNode.next = newNode
		newNode.prev = newNode
		l.head = newNode
		return
	}
	tail := l.head.prev
	tail.next = newNode
	newNode.prev = tail
	newNode.next = l.head
	l.head.prev = newNode
}

func (l *CircularList) Delete(value int) {
	if l.head == nil {
		return
	}
	cur := l.head
	for {
		if cur.value == value {
			if cur.next == cur {
				l.head = nil
				return
			}
			cur.prev.next = cur.next
			cur.next.prev = cur.prev
			if cur == l.head {
				l.head = cur.next
			}
			return
		}
		cur = cur.next
		if cur == l.head {
			break
		}
	}
}

func (l *CircularList) Print() {
	if l.head == nil {
		fmt.Println("(пусто)")
		return
	}
	cur := l.head
	for {
		fmt.Printf("%d -> ", cur.value)
		cur = cur.next
		if cur == l.head {
			break
		}
	}
	fmt.Println("(head)")
}

func main() {
	l := &CircularList{}
	for _, v := range []int{1, 2, 3, 4, 5} {
		l.Add(v)
	}
	fmt.Print("Список: ")
	l.Print()

	l.Delete(3)
	fmt.Print("После удаления 3: ")
	l.Print()

	l.Delete(1)
	fmt.Print("После удаления 1 (head): ")
	l.Print()

	l.Delete(5)
	fmt.Print("После удаления 5 (tail): ")
	l.Print()

	l2 := &CircularList{}
	l2.Add(42)
	fmt.Print("Список из одного элемента: ")
	l2.Print()
	l2.Delete(42)
	fmt.Print("После удаления единственного: ")
	l2.Print()
}
