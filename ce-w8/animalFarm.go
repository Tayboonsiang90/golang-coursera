package main

import (
	"fmt"
)

type Animal interface {
	Eat()
	Move()
	Speak()
	GetName() string
}

type Cow struct{ Name string }

func (c Cow) Eat()            { fmt.Println("grass") }
func (c Cow) Move()           { fmt.Println("walk") }
func (c Cow) Speak()          { fmt.Println("moo") }
func (c Cow) GetName() string { return c.Name }

type Bird struct{ Name string }

func (b Bird) Eat()            { fmt.Println("worms") }
func (b Bird) Move()           { fmt.Println("fly") }
func (b Bird) Speak()          { fmt.Println("peep") }
func (b Bird) GetName() string { return b.Name }

type Snake struct{ Name string }

func (s Snake) Eat()            { fmt.Println("mice") }
func (s Snake) Move()           { fmt.Println("slither") }
func (s Snake) Speak()          { fmt.Println("hsss") }
func (s Snake) GetName() string { return s.Name }

func main() {
	sli := make([]Animal, 0, 0)
	for {
		fmt.Printf("> ")
		var query, name, queryDetails string
		fmt.Scanf("%s %s %s", &query, &name, &queryDetails)
		var newAnimal Animal
		if query == "newanimal" {
			switch queryDetails {
			case "cow":
				newAnimal = Cow{name}
			case "bird":
				newAnimal = Bird{name}
			case "snake":
				newAnimal = Snake{name}
			}
			sli = append(sli, newAnimal)
			fmt.Println("Created it!")
		} else if query == "query" {
			for _, v := range sli {
				if v.GetName() == name {
					switch queryDetails {
					case "eat":
						v.Eat()
					case "move":
						v.Move()
					case "speak":
						v.Speak()
					}
				}
			}
		}
	}
}
