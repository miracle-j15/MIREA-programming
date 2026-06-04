package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	issued := make([]int, n)
	for i := range issued {
		fmt.Scan(&issued[i])
	}

	var m int
	fmt.Scan(&m)
	returned := make(map[int]bool, m)
	for i := 0; i < m; i++ {
		var x int
		fmt.Scan(&x)
		returned[x] = true
	}

	for _, badge := range issued {
		if !returned[badge] {
			fmt.Println(badge)
		}
	}
}
