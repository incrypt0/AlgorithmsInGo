package main

import "fmt"

func main() {
	arr := [10]int{7, 3, 8, 5, 1, 9, 2, 4, 0, 6}
	for i := 0; i < 10; i++ {
		for j := 0; j < 9; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	fmt.Println(arr)

}
