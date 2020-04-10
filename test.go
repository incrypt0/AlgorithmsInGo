package main

import "fmt"
func change(a *[]int){
	fmt.Println((*a)[0])
}
func main() {
	a:= []int{1,2,3,4,5,6,7}
	change(&a)
	fmt.Println(a)
}