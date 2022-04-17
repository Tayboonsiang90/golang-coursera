package main

import (
	// "fmt" // no more error
)

func main() {
	//structs
	//can use different types

	type Person struct {
		name string
		addr string
		phone int
	}
	//initialize 1
	var p1 Person
	p1.name="Joe"
	p1.addr="bukit batok"
	p1.phone=12

	//initialize 2
	p2 := new(Person)

	//initialize 3
	p3 := Person{name: "joe", addr: "bukit batok", phone: 123}

	_ = p2
	_ = p3
}
