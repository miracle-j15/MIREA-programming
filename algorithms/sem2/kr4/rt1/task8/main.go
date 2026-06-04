package main

import (
	"fmt"
	"math/rand"
	"time"
)

var cmp, swp int64

func heapifyDown(arr []int, n, i int) {
	largest, left, right := i, 2*i+1, 2*i+2
	if left < n {
		cmp++
		if arr[left] > arr[largest] {
			largest = left
		}
	}
	if right < n {
		cmp++
		if arr[right] > arr[largest] {
			largest = right
		}
	}
	if largest != i {
		arr[i], arr[largest] = arr[largest], arr[i]
		swp++
		heapifyDown(arr, n, largest)
	}
}

func heapSort(arr []int) {
	n := len(arr)
	for i := n/2 - 1; i >= 0; i-- {
		heapifyDown(arr, n, i)
	}
	for i := n - 1; i > 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		swp++
		heapifyDown(arr, i, 0)
	}
}

func bubbleSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			cmp++
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				swp++
			}
		}
	}
}

func insertionSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		key, j := arr[i], i-1
		for j >= 0 {
			cmp++
			if arr[j] <= key {
				break
			}
			arr[j+1] = arr[j]
			swp++
			j--
		}
		arr[j+1] = key
	}
}

func selectionSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		minIdx := i
		for j := i + 1; j < n; j++ {
			cmp++
			if arr[j] < arr[minIdx] {
				minIdx = j
			}
		}
		if minIdx != i {
			arr[i], arr[minIdx] = arr[minIdx], arr[i]
			swp++
		}
	}
}

func quickSort(arr []int, lo, hi int) {
	if lo >= hi {
		return
	}
	pivot, i := arr[hi], lo-1
	for j := lo; j < hi; j++ {
		cmp++
		if arr[j] <= pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
			swp++
		}
	}
	arr[i+1], arr[hi] = arr[hi], arr[i+1]
	swp++
	quickSort(arr, lo, i)
	quickSort(arr, i+2, hi)
}

func bench(name string, fn func([]int), base []int) {
	arr := append([]int{}, base...)
	cmp, swp = 0, 0
	start := time.Now()
	fn(arr)
	elapsed := time.Since(start)
	fmt.Printf("  %-16s time=%-10s cmp=%-10d swaps=%d\n", name, elapsed.Round(time.Microsecond), cmp, swp)
}

func benchQS(base []int) {
	arr := append([]int{}, base...)
	cmp, swp = 0, 0
	start := time.Now()
	quickSort(arr, 0, len(arr)-1)
	elapsed := time.Since(start)
	fmt.Printf("  %-16s time=%-10s cmp=%-10d swaps=%d\n", "QuickSort", elapsed.Round(time.Microsecond), cmp, swp)
}

func runTest(desc string, base []int) {
	fmt.Printf("\n%s (n=%d):\n", desc, len(base))
	bench("HeapSort", heapSort, base)
	bench("BubbleSort", bubbleSort, base)
	bench("InsertionSort", insertionSort, base)
	bench("SelectionSort", selectionSort, base)
	benchQS(base)
}

func main() {
	runTest("Малый частично отсортированный", []int{1, 2, 4, 3, 5, 7, 6, 8, 10, 9})
	runTest("Малый по убыванию", []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1})
	runTest("Малый по возрастанию", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	runTest("Малый случайный", []int{5, 3, 8, 1, 9, 2, 7, 4, 6, 10})

	rng := rand.New(rand.NewSource(42))
	large := make([]int, 10000)
	for i := range large {
		large[i] = rng.Intn(100000)
	}
	runTest("Большой случайный", large)

	fmt.Println("\nВывод:")
	fmt.Println("  HeapSort:      O(n log n) всегда, стабильно на любых данных")
	fmt.Println("  QuickSort:     O(n log n) в среднем, деградирует на отсортированных")
	fmt.Println("  InsertionSort: O(n) на почти отсортированных, O(n²) в худшем")
	fmt.Println("  BubbleSort:    O(n²) всегда, максимум перестановок")
	fmt.Println("  SelectionSort: O(n²) всегда, минимум перестановок")
}
