package main

import (
	"fmt"
)

type Animal struct {
	food       string
	locomotion string
	noise      string
}

func (ani Animal) Eat() {
	println(ani.food)
}
func (ani Animal) Move() {
	println(ani.locomotion)
}
func (ani Animal) Speak() {
	println(ani.noise)
}

func checkAnimal(animal, method string) {
	cow := Animal{"grass", "walk", "moo"}
	bird := Animal{"worms", "fly", "peep"}
	snake := Animal{"mice", "slither", "hsss"}
	if animal == "cow" {
		checkMethod(cow, method)
	} else if animal == "bird" {
		checkMethod(bird, method)
	} else if animal == "snake" {
		checkMethod(snake, method)
	}
}

func checkMethod(animal Animal, method string) {
	if method == "eat" {
		animal.Eat()
	} else if method == "move" {
		animal.Move()
	} else if method == "speak" {
		animal.Speak()
	}
}

func main() {
	fmt.Printf("> ")
	var animal string
	var method string
	fmt.Scanf("%s %s", &animal, &method)

	checkAnimal(animal, method)
}
