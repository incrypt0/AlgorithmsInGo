package main

import "fmt"

func main() {
	num := 10
	//in go for is itself the keyword used for while loop
	for num > 0 {
		fmt.Println(num)
		num--
	}
}
