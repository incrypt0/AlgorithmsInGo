package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//Taking User Inputs
	var name string
	// //Scan method inside fmt
	// fmt.Println("Hey there, Enter your name ?")
	// fmt.Scan(&name)
	// fmt.Println(name)

	// //Scanln
	// fmt.Println("Hey there, Enter your name ?")
	// fmt.Scanln(&name)
	// fmt.Println(name)

	//Using bufio package
	//ReadLine method
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Hey there, Enter your Full Name : ")
	blah, _, _ := reader.ReadLine()
	fmt.Println(string(blah))

	//Using the ReadString method
	fmt.Print("Hey there, Enter your Full Name : ")
	name, _ = reader.ReadString('\n')
	fmt.Println(name)
	fmt.Println('a')

	//Scanner in bufio
}
