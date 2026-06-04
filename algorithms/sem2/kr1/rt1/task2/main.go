package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	arr := make([]int, n)
	total := 0
	for i := range arr {
		fmt.Scan(&arr[i])
		total += arr[i]
	}

	if total%2 != 0 {
		fmt.Println("NO")
		return
	}

	half := total / 2
	prefix := 0
	for _, v := range arr {
		prefix += v
		if prefix == half {
			fmt.Println("YES")
			return
		}
	}
	fmt.Println("NO")
}
