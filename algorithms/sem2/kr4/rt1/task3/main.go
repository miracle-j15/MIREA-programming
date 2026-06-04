package main

import (
	"errors"
	"fmt"
)

type BinaryHeap struct {
	data []int
}

func (h *BinaryHeap) heapifyUp(i int) {
	for i > 0 {
		parent := (i - 1) / 2
		if h.data[parent] < h.data[i] {
			h.data[parent], h.data[i] = h.data[i], h.data[parent]
			i = parent
		} else {
			break
		}
	}
}

func (h *BinaryHeap) heapifyDown(i int) {
	n := len(h.data)
	for {
		largest := i
		left, right := 2*i+1, 2*i+2
		if left < n && h.data[left] > h.data[largest] {
			largest = left
		}
		if right < n && h.data[right] > h.data[largest] {
			largest = right
		}
		if largest == i {
			break
		}
		h.data[i], h.data[largest] = h.data[largest], h.data[i]
		i = largest
	}
}

func NewBinaryHeap(arr []int) *BinaryHeap {
	h := &BinaryHeap{data: append([]int{}, arr...)}
	for i := len(h.data)/2 - 1; i >= 0; i-- {
		h.heapifyDown(i)
	}
	return h
}

func (h *BinaryHeap) GetMax() (int, error) {
	if len(h.data) == 0 {
		return 0, errors.New("куча пуста")
	}
	return h.data[0], nil
}

func (h *BinaryHeap) ExtractMax() (int, error) {
	if len(h.data) == 0 {
		return 0, errors.New("куча пуста")
	}
	max := h.data[0]
	last := len(h.data) - 1
	h.data[0] = h.data[last]
	h.data = h.data[:last]
	if len(h.data) > 0 {
		h.heapifyDown(0)
	}
	return max, nil
}

func (h *BinaryHeap) Insert(value int) {
	h.data = append(h.data, value)
	h.heapifyUp(len(h.data) - 1)
}

func (h *BinaryHeap) Empty() bool { return len(h.data) == 0 }
func (h *BinaryHeap) Size() int   { return len(h.data) }
func (h *BinaryHeap) Print()      { fmt.Println(h.data) }

func main() {
	heap := NewBinaryHeap([]int{3, 1, 6, 5, 2, 4})
	fmt.Print("Куча из {3,1,6,5,2,4}: ")
	heap.Print()

	max, _ := heap.GetMax()
	fmt.Println("GetMax:", max)

	heap.Insert(8)
	fmt.Print("После Insert(8): ")
	heap.Print()

	extracted, _ := heap.ExtractMax()
	fmt.Println("ExtractMax:", extracted)
	fmt.Print("Куча после извлечения: ")
	heap.Print()

	empty := &BinaryHeap{}
	empty.Insert(5)
	empty.Insert(10)
	empty.Insert(3)
	fmt.Print("Пустая куча + Insert(5,10,3): ")
	empty.Print()
}
