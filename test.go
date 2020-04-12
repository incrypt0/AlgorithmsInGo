package main

import (
	"fmt"
	"strings"
)

type Animal struct {
	name  string
	food  string
	move  string
	sound string
}

//methods
func (ani Animal) Eat() {
	fmt.Println(ani.food)
}
func (ani Animal) Move() {
	fmt.Println(ani.move)
}
func (ani Animal) Speak() {
	fmt.Println(ani.sound)
}
func ProcessRes(animalName string, featureName string, animalMap map[string]Animal) {
	animNameLower := strings.ToLower(animalName)
	featureNameLower := strings.ToLower(featureName)
	var ani Animal
	switch animNameLower {
	case "cow":
		ani = animalMap["cow"]
	case "snake":
		ani = animalMap["snake"]
	case "bird":
		ani = animalMap["snake"]
	}
	switch featureNameLower {
	case "eat":
		ani.Eat()
	case "move":
		ani.Move()
	case "speak":
		ani.Speak()
	}
}
func main() {
	var a, b string

	cow := Animal{
		name:  "cow",
		food:  "grass",
		move:  "walk",
		sound: "moo",
	}
	bird := Animal{
		name:  "bird",
		food:  "worms",
		move:  "fly",
		sound: "peep",
	}
	snake := Animal{
		name:  "snake",
		food:  "mice",
		move:  "slither",
		sound: "hsss",
	}
	animalMap := map[string]Animal{
		"cow":   cow,
		"bird":  bird,
		"snake": snake,
	}
	for {
		fmt.Print("> ")
		fmt.Scanf("%s %s", &a, &b)
		ProcessRes(a, b, animalMap)
	}
}
