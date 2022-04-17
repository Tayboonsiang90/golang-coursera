package main

import (
	"fmt"
)

//reciever type
type MyInt int
type Coordinates struct {
	X    float64
	Y    float64
	Dist float64
}

func (mi MyInt) Double() int {
	return int(mi * 2)
}

func (coords *Coordinates) Distance() {
	coords.Dist = coords.X + coords.Y
}
func (coords *Coordinates) Increment() {
	coords.X++
	coords.Y++
}

func main() {
	v := MyInt(3)
	//The implicit argument is v
	fmt.Println(v.Double())

	//My own learning
	coords := Coordinates{1, 2, 0}
	coords.Increment()
	coords.Distance()
	fmt.Println(coords)

}
