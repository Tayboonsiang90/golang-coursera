package main

import (
	"fmt"
)

type Speaker interface {
	Speak()
}

type Dog struct {
	name string
}

func (d *Dog) Speak() {
	//should check if there is dynamic value or not
	if (d == nil) {
		fmt.Println("<noise>")
	} else {
		fmt.Println(d.name)
	}
}

func main() {
	var s1 Speaker
	d1 := Dog{"Brian"}
	//an interface has a dynamic type and a dynamic value
	//the dynamic type is dog and the dynamic value is d1
	d1.Speak()
	s1 = &d1
	s1.Speak()

	//interfaces with nil dynamic value but has a dynamic type
	//s2 has a dynamic type but no dynamic value
	var s2 Speaker
	//d2 has no concrete value
	var d2 *Dog
	s2 = d2
	s2.Speak()
}
