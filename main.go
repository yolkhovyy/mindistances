package main

import (
	"fmt"
)

func minDistance(a *[]int) int {
	min := len(*a)
	for i := 0; i < len(*a); i++ {
		for j := i + 1; j < len(*a); j++ {
			if (*a)[i] == (*a)[j] {
				if j-i < min {
					min = j - i
				}
			}
		}
	}
	if min == len(*a) {
		min = -1
	}
	return min
}

func main() {
	a := []int{3, 2, 0, 8, 2, 3}
	fmt.Printf("%d", minDistance(&a))
}
