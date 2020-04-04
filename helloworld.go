package main

import "fmt"

func add(num1, num2 float64) float64 { //This is same ass (num1 float4 ,num2 float64)
	return num1 + num2
}
func main() {
	fmt.Println("Hello,World!\nHere is an addition")
	fmt.Println(add(2.4, 3.2))
}
