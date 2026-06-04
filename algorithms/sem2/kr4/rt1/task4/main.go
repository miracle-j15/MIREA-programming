package main

import (
	"container/heap"
	"fmt"
)

type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x any)        { *h = append(*h, x.(int)) }
func (h *MinHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

type Step struct{ a, b int }

func main() {
	var n int
	fmt.Scan(&n)

	h := &MinHeap{}
	heap.Init(h)
	for i := 0; i < n; i++ {
		var x int
		fmt.Scan(&x)
		heap.Push(h, x)
	}

	var totalCost int
	var steps []Step

	for h.Len() > 1 {
		a := heap.Pop(h).(int)
		b := heap.Pop(h).(int)
		combined := a + b
		totalCost += combined
		steps = append(steps, Step{a, b})
		heap.Push(h, combined)
	}

	fmt.Println("Порядок связывания:")
	for _, s := range steps {
		fmt.Printf("%d + %d = %d\n", s.a, s.b, s.a+s.b)
	}
	fmt.Println("Суммарные затраты:", totalCost)
}
