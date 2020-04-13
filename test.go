package main

import (
	"fmt"
	"strings"
)

type Animal interface {
	Eat()
	Speak()
	Move()
	ChangeName(string)
}
type cow struct {
	name string
}

func (c cow) Eat() {
	fmt.Println("grass")
}
func (c cow) Move() {
	fmt.Println("walk")
}
func (c cow) Speak() {
	fmt.Println("moo")
}
func (c cow) ChangeName(name string) {
	c.name = name
}

type bird struct {
	name string
}

func (c bird) Eat() {
	fmt.Println("worm")
}
func (c bird) Move() {
	fmt.Println("fly")
}
func (c bird) Speak() {
	fmt.Println("peep")
}
func (c bird) ChangeName(name string) {
	c.name = name
}

type snake struct {
	name string
}

func (c snake) Eat() {
	fmt.Println("mice")
}
func (c snake) Move() {
	fmt.Println("slither")
}
func (c snake) Speak() {
	fmt.Println("hss")
}
func (c snake) ChangeName(name string) {
	c.name = name
}
func main() {
	var a, b, c string
	animMap := make(map[string]Animal)
	animTypeMap := map[string]Animal{
		"cow":   cow{},
		"snake": snake{},
		"bird":  bird{},
	}

	for {
		fmt.Print("> ")
		fmt.Scanf("%s %s %s", &a, &b, &c)
		scanned := strings.ToLower(a)
		if scanned == "newanimal" {
			anim := animTypeMap[strings.ToLower(c)]
			anim.ChangeName(b)
			animMap[strings.ToLower(b)] = anim
			fmt.Println("Created..!!")
		}

		if scanned == "query" {
			item := animMap[strings.ToLower(b)]
			command := strings.ToLower(c)
			switch command {
			case "move":
				item.Move()
			case "eat":
				item.Eat()
			case "speak":
				item.Speak()
			}
		}

	}
}
