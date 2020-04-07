package main

import "fmt"

func foo() *int {
	x := 2
	return &x
}
func main() {
	y := '1'
	x := int(y)
	fmt.Println(x)
}
