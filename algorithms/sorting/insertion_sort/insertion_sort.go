package main

import "fmt"

func insertionSort() {

}
func main() {
	arr := [10]int{7, 3, 8, 5, 1, 9, 2, 4, 0, 6}

	for i := 0; i < len(arr); i++ {
		fmt.Println(i, arr[i])
		item1 := arr[i]
		j := i - 1

		for j >= 0 && item1 < arr[j] {

			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = item1
	}
	fmt.Println(arr)
}
