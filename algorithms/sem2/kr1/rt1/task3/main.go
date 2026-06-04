package main

import "fmt"

func majorityElement(arr []int) (int, bool) {
	candidate, count := 0, 0
	for _, v := range arr {
		if count == 0 {
			candidate = v
		}
		if v == candidate {
			count++
		} else {
			count--
		}
	}
	cnt := 0
	for _, v := range arr {
		if v == candidate {
			cnt++
		}
	}
	if cnt > len(arr)/2 {
		return candidate, true
	}
	return 0, false
}

func main() {
	var n int
	fmt.Scan(&n)
	arr := make([]int, n)
	for i := range arr {
		fmt.Scan(&arr[i])
	}

	if leader, ok := majorityElement(arr); ok {
		fmt.Println(leader)
	} else {
		fmt.Println("Нет лидера")
	}
}
