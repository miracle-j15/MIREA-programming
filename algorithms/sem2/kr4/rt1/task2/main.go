package main

import "fmt"

type Homework struct {
	student    string
	submitTime int
}

type HomeworkQueue struct {
	heap []Homework
}

func (q *HomeworkQueue) heapifyUp(i int) {
	for i > 0 {
		parent := (i - 1) / 2
		if q.heap[parent].submitTime > q.heap[i].submitTime {
			q.heap[parent], q.heap[i] = q.heap[i], q.heap[parent]
			i = parent
		} else {
			break
		}
	}
}

func (q *HomeworkQueue) heapifyDown(i int) {
	n := len(q.heap)
	for {
		smallest := i
		left, right := 2*i+1, 2*i+2
		if left < n && q.heap[left].submitTime < q.heap[smallest].submitTime {
			smallest = left
		}
		if right < n && q.heap[right].submitTime < q.heap[smallest].submitTime {
			smallest = right
		}
		if smallest == i {
			break
		}
		q.heap[i], q.heap[smallest] = q.heap[smallest], q.heap[i]
		i = smallest
	}
}

func (q *HomeworkQueue) Insert(student string, time int) {
	q.heap = append(q.heap, Homework{student, time})
	q.heapifyUp(len(q.heap) - 1)
}

func (q *HomeworkQueue) ExtractMin() Homework {
	top := q.heap[0]
	last := len(q.heap) - 1
	q.heap[0] = q.heap[last]
	q.heap = q.heap[:last]
	if len(q.heap) > 0 {
		q.heapifyDown(0)
	}
	return top
}

func (q *HomeworkQueue) Empty() bool { return len(q.heap) == 0 }

func main() {
	var n int
	fmt.Scan(&n)

	q := &HomeworkQueue{}
	for i := 0; i < n; i++ {
		var student string
		var t int
		fmt.Scan(&student, &t)
		q.Insert(student, t)
	}

	fmt.Println("Порядок проверки:")
	for !q.Empty() {
		hw := q.ExtractMin()
		fmt.Printf("%s (время: %d)\n", hw.student, hw.submitTime)
	}
}
