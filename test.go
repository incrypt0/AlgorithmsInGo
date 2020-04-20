package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	sli := []int{1, 5, 8, 6, 7, 3, 4, 87, 10}
	a, b, c, d := split4(sli)
	fmt.Println(a, b, c, d)

	wg.Add(4)
	go func() {
		Sort(&a)
		wg.Done()
	}()
	go func() {
		Sort(&b)
		wg.Done()
	}()
	go func() {
		Sort(&c)
		wg.Done()
	}()
	go func() {
		Sort(&d)
		wg.Done()
	}()
	wg.Wait()
	res := merge(merge(a, b), merge(c, d))
	fmt.Println(res)
}
func merge(left, right []int) (result []int) {
	result = make([]int, len(left)+len(right))

	i := 0
	for len(left) > 0 && len(right) > 0 {
		if left[0] < right[0] {
			result[i] = left[0]
			left = left[1:]
		} else {
			result[i] = right[0]
			right = right[1:]
		}
		i++
	}

	for j := 0; j < len(left); j++ {
		result[i] = left[j]
		i++
	}
	for j := 0; j < len(right); j++ {
		result[i] = right[j]
		i++
	}

	return result
}
func split4(sli []int) ([]int, []int, []int, []int) {
	len := len(sli)
	lenPiece := len / 4
	a, b, c, d := sli[0:lenPiece], sli[lenPiece:2*lenPiece], sli[2*lenPiece:3*lenPiece], sli[3*lenPiece:]
	return a, b, c, d
}
func swap(num1 *int, num2 *int) {
	*num1, *num2 = *num2, *num1
}
func Sort(arrPtr *[]int) {
	arr := *arrPtr
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-1; j++ {
			if arr[j] > arr[j+1] {
				swap(&arr[j], &arr[j+1])
			}
		}
	}

}
