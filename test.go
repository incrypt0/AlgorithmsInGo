package main

import (
	"fmt"
)

func sortItBaby(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		item1 := arr[i]
		j := i - 1
		for j >= 0 && item1 < arr[j] {
			arr[j+1], arr[j] = arr[j], arr[j+1]
			j--
		}
	}
	return arr
}

type P struct {
	x string
	y int
}

func main() {
	s := make([]int, 0, 3)
	s = append(s, 100)
	fmt.Println(len(s), cap(s))
}
