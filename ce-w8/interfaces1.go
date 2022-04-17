package main

import (
	"fmt"
)

type Shape2D interface {
	Area()
	Perimeter()
}

type Triangle struct {
	base   float64
	height float64
}

func (t Triangle) Area() {
	fmt.Printf("The area is %f", 0.5*t.base*t.height)
}
func (t Triangle) Perimeter() {
	fmt.Printf("The perimeter is %f %f", t.base, t.height)
}

func main() {
	t1 := Triangle{1,2}
	var s1 Shape2D
	s1 = t1
	s1.Area()
	s1.Perimeter()
	fmt.Println(t1)
}
