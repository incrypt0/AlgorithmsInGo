package main

func lowestIndex(arr []int) int {
	var length, minIndex int = len(arr), 0
	for i := 0; i < length; i++ {
		if arr[i] < arr[minIndex] {
			minIndex = i
		}
	}
	return minIndex
}
func main() {
	arr := [10]int{7, 3, 8, 5, 1, 9, 2, 4, 0, 6}
	length := len(arr)
	for i := 0; i < length; i++ {
		lowestIndex(arr[i+1:])
	}
}
