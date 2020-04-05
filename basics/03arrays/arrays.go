package main

import (
	"fmt"
)

func main() {
	var arr1 [4]int
	var arr2 = [4]int{1, 2, 3, 4}
	var arr3 = [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr1, arr2, arr3)
	fmt.Println(arr3[:3])
}
