package main

import (
	"errors"
	"fmt"
)

type Node struct {
	data     int
	priority int
	next     *Node
}

type PriorityQueue struct {
	head *Node
}

func (q *PriorityQueue) Insert(data, priority int) {
	newNode := &Node{data: data, priority: priority}
	if q.head == nil || priority > q.head.priority {
		newNode.next = q.head
		q.head = newNode
		return
	}
	cur := q.head
	for cur.next != nil && cur.next.priority >= priority {
		cur = cur.next
	}
	newNode.next = cur.next
	cur.next = newNode
}

func (q *PriorityQueue) ExtractMax() (int, error) {
	if q.head == nil {
		return 0, errors.New("очередь пуста")
	}
	val := q.head.data
	q.head = q.head.next
	return val, nil
}

func (q *PriorityQueue) Empty() bool { return q.head == nil }

func (q *PriorityQueue) Print() {
	for cur := q.head; cur != nil; cur = cur.next {
		fmt.Printf("[data=%d pri=%d] -> ", cur.data, cur.priority)
	}
	fmt.Println("nil")
}

func main() {
	q := &PriorityQueue{}
	q.Insert(10, 1)
	q.Insert(20, 3)
	q.Insert(30, 2)
	q.Insert(40, 5)
	q.Insert(50, 4)

	fmt.Print("Очередь: ")
	q.Print()

	fmt.Println("Извлечение по убыванию приоритета:")
	for !q.Empty() {
		val, _ := q.ExtractMax()
		fmt.Printf("  data=%d\n", val)
	}
}
