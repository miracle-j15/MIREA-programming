package main

import "fmt"

func heapifyDown(arr []int, n, i int) {
	largest := i
	left, right := 2*i+1, 2*i+2
	if left < n && arr[left] > arr[largest] {
		largest = left
	}
	if right < n && arr[right] > arr[largest] {
		largest = right
	}
	if largest != i {
		arr[i], arr[largest] = arr[largest], arr[i]
		heapifyDown(arr, n, largest)
	}
}

func buildMaxHeap(arr []int) {
	n := len(arr)
	for i := n/2 - 1; i >= 0; i-- {
		heapifyDown(arr, n, i)
	}
}

func main() {
	var n int
	fmt.Scan(&n)
	arr := make([]int, n)
	for i := range arr {
		fmt.Scan(&arr[i])
	}

	buildMaxHeap(arr)

	for i, v := range arr {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(v)
	}
	fmt.Println()
}
