package main

import (
	"fmt"
	"sort"
)

func assignCouriers(orders []int, maxRoute int) (int, [][]int) {
	sorted := append([]int{}, orders...)
	sort.Sort(sort.Reverse(sort.IntSlice(sorted)))

	var routes [][]int
	var loads []int

	for _, dist := range sorted {
		if dist > maxRoute {
			fmt.Printf("  [!] Заказ %dкм превышает лимит %dкм, пропущен\n", dist, maxRoute)
			continue
		}
		assigned := false
		best := -1
		for i, load := range loads {
			if load+dist <= maxRoute {
				if best == -1 || loads[i] < loads[best] {
					best = i
				}
			}
		}
		if best != -1 {
			loads[best] += dist
			routes[best] = append(routes[best], dist)
			assigned = true
		}
		if !assigned {
			loads = append(loads, dist)
			routes = append(routes, []int{dist})
		}
	}

	return len(routes), routes
}

func main() {
	orders := []int{10, 5, 8, 3, 7, 12, 2, 9}
	maxRoute := 20

	fmt.Printf("Заказы: %v\n", orders)
	fmt.Printf("Лимит маршрута: %dкм\n\n", maxRoute)

	count, routes := assignCouriers(orders, maxRoute)

	for i, route := range routes {
		total := 0
		for _, d := range route {
			total += d
		}
		fmt.Printf("Курьер %d: заказы %v, итого %dкм\n", i+1, route, total)
	}
	fmt.Printf("\nЗадействовано курьеров: %d\n", count)
}
