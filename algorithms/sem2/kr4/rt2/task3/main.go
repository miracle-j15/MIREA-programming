package main

import (
	"container/heap"
	"fmt"
	"sort"
)

type Exam struct {
	start, end int
}

type EndTimeHeap []int

func (h EndTimeHeap) Len() int           { return len(h) }
func (h EndTimeHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h EndTimeHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *EndTimeHeap) Push(x any)        { *h = append(*h, x.(int)) }
func (h *EndTimeHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func minRooms(exams []Exam) int {
	sort.Slice(exams, func(i, j int) bool {
		return exams[i].start < exams[j].start
	})

	rooms := &EndTimeHeap{}
	heap.Init(rooms)

	for _, exam := range exams {
		if rooms.Len() > 0 && (*rooms)[0] <= exam.start {
			heap.Pop(rooms)
		}
		heap.Push(rooms, exam.end)
	}

	return rooms.Len()
}

func main() {
	exams := []Exam{
		{1, 4},
		{2, 5},
		{6, 8},
	}

	fmt.Println("Расписание экзаменов:")
	for _, e := range exams {
		fmt.Printf("  [%d, %d]\n", e.start, e.end)
	}
	fmt.Println("Минимум аудиторий:", minRooms(exams))

	exams2 := []Exam{{1, 10}, {2, 3}, {4, 5}, {6, 7}}
	fmt.Println("\nРасписание 2:")
	for _, e := range exams2 {
		fmt.Printf("  [%d, %d]\n", e.start, e.end)
	}
	fmt.Println("Минимум аудиторий:", minRooms(exams2))
}
