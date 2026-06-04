package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	arr := make([]int, n)
	for i := range arr {
		fmt.Scan(&arr[i])
	}

	maxLen, cur := 1, 1
	for i := 1; i < n; i++ {
		if arr[i] == arr[i-1] {
			cur++
			if cur > maxLen {
				maxLen = cur
			}
		} else {
			cur = 1
		}
	}

	fmt.Println(maxLen)
}
