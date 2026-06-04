package main

import (
	"errors"
	"fmt"
	"sort"
)

type Task struct {
	name     string
	priority int
}

type PriorityQueue struct {
	heap []Task
}

func (q *PriorityQueue) heapifyUp(i int) {
	for i > 0 {
		parent := (i - 1) / 2
		if q.heap[parent].priority < q.heap[i].priority {
			q.heap[parent], q.heap[i] = q.heap[i], q.heap[parent]
			i = parent
		} else {
			break
		}
	}
}

func (q *PriorityQueue) heapifyDown(i int) {
	n := len(q.heap)
	for {
		largest := i
		left, right := 2*i+1, 2*i+2
		if left < n && q.heap[left].priority > q.heap[largest].priority {
			largest = left
		}
		if right < n && q.heap[right].priority > q.heap[largest].priority {
			largest = right
		}
		if largest == i {
			break
		}
		q.heap[i], q.heap[largest] = q.heap[largest], q.heap[i]
		i = largest
	}
}

func (q *PriorityQueue) Insert(name string, priority int) {
	q.heap = append(q.heap, Task{name, priority})
	q.heapifyUp(len(q.heap) - 1)
}

func (q *PriorityQueue) ExtractMax() (Task, error) {
	if len(q.heap) == 0 {
		return Task{}, errors.New("список пуст")
	}
	top := q.heap[0]
	last := len(q.heap) - 1
	q.heap[0] = q.heap[last]
	q.heap = q.heap[:last]
	if len(q.heap) > 0 {
		q.heapifyDown(0)
	}
	return top, nil
}

func (q *PriorityQueue) Edit(name, newName string, newPriority int) bool {
	for i := range q.heap {
		if q.heap[i].name == name {
			q.heap[i].name = newName
			q.heap[i].priority = newPriority
			q.heapifyUp(i)
			q.heapifyDown(i)
			return true
		}
	}
	return false
}

func (q *PriorityQueue) Empty() bool { return len(q.heap) == 0 }

func (q *PriorityQueue) PrintAll() {
	if q.Empty() {
		fmt.Println("  (список пуст)")
		return
	}
	tmp := append([]Task{}, q.heap...)
	sort.Slice(tmp, func(i, j int) bool { return tmp[i].priority > tmp[j].priority })
	for _, t := range tmp {
		fmt.Printf("  [%d] %s\n", t.priority, t.name)
	}
}

func main() {
	q := &PriorityQueue{}
	fmt.Println("=== TODO LIST ===")

	for {
		fmt.Print("\n1. Добавить  2. Выполнить  3. Редактировать  4. Показать  0. Выход\n> ")
		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 0:
			return
		case 1:
			var name string
			var priority int
			fmt.Print("Название: ")
			fmt.Scan(&name)
			fmt.Print("Приоритет: ")
			fmt.Scan(&priority)
			q.Insert(name, priority)
			fmt.Println("Добавлено.")
		case 2:
			t, err := q.ExtractMax()
			if err != nil {
				fmt.Println("Список пуст.")
			} else {
				fmt.Printf("Выполнено: [%d] %s\n", t.priority, t.name)
			}
		case 3:
			var name, newName string
			var newPriority int
			fmt.Print("Название задачи: ")
			fmt.Scan(&name)
			fmt.Print("Новое название: ")
			fmt.Scan(&newName)
			fmt.Print("Новый приоритет: ")
			fmt.Scan(&newPriority)
			if q.Edit(name, newName, newPriority) {
				fmt.Println("Обновлено.")
			} else {
				fmt.Println("Задача не найдена.")
			}
		case 4:
			fmt.Println("Задачи (по приоритету):")
			q.PrintAll()
		}
	}
}
