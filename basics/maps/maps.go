package main

import "fmt"

func main() {
	var idMap1 map[string]int
	idMap1 = make(map[string]int)
	idMap2 := map[string]int{
		"RollNo": 10,
	}
	fmt.Println(idMap1, idMap2)
}


