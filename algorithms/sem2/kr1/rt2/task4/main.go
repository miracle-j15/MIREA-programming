package main

import "fmt"

type Node struct {
	value int
	prev  *Node
	next  *Node
}

type DoublyList struct {
	head *Node
	tail *Node
	size int
}

func (l *DoublyList) Add(value int) {
	n := &Node{value: value}
	if l.tail == nil {
		l.head = n
		l.tail = n
	} else {
		n.prev = l.tail
		l.tail.next = n
		l.tail = n
	}
	l.size++
}

func (l *DoublyList) InsertBefore(target, value int) {
	for cur := l.head; cur != nil; cur = cur.next {
		if cur.value == target {
			n := &Node{value: value, prev: cur.prev, next: cur}
			if cur.prev != nil {
				cur.prev.next = n
			} else {
				l.head = n
			}
			cur.prev = n
			l.size++
			return
		}
	}
}

func (l *DoublyList) DeleteAll(value int) {
	for cur := l.head; cur != nil; {
		next := cur.next
		if cur.value == value {
			if cur.prev != nil {
				cur.prev.next = cur.next
			} else {
				l.head = cur.next
			}
			if cur.next != nil {
				cur.next.prev = cur.prev
			} else {
				l.tail = cur.prev
			}
			l.size--
		}
		cur = next
	}
}

func (l *DoublyList) Count() int { return l.size }

func (l *DoublyList) PrintForward() {
	for cur := l.head; cur != nil; cur = cur.next {
		fmt.Printf("%d ", cur.value)
	}
	fmt.Println()
}

func (l *DoublyList) PrintBackward() {
	for cur := l.tail; cur != nil; cur = cur.prev {
		fmt.Printf("%d ", cur.value)
	}
	fmt.Println()
}

func (l *DoublyList) IsPalindrome() bool {
	left, right := l.head, l.tail
	for left != nil && right != nil && left != right && left.prev != right {
		if left.value != right.value {
			return false
		}
		left = left.next
		right = right.prev
	}
	return true
}

func main() {
	l := &DoublyList{}
	for _, v := range []int{1, 2, 3, 2, 1} {
		l.Add(v)
	}
	fmt.Print("Прямой:   ")
	l.PrintForward()
	fmt.Print("Обратный: ")
	l.PrintBackward()
	fmt.Println("Палиндром:", l.IsPalindrome())
	fmt.Println("Размер:", l.Count())

	l.InsertBefore(3, 99)
	fmt.Print("После InsertBefore(3, 99): ")
	l.PrintForward()

	l.DeleteAll(2)
	fmt.Print("После DeleteAll(2): ")
	l.PrintForward()
	fmt.Println("Палиндром:", l.IsPalindrome())

	l2 := &DoublyList{}
	l2.Add(1)
	fmt.Println("Один элемент — палиндром:", l2.IsPalindrome())

	empty := &DoublyList{}
	fmt.Println("Пустой — палиндром:", empty.IsPalindrome())
}
