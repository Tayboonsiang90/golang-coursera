package main

import (
	"fmt"
)

func printHello(x, y string) (string, int) {
	fmt.Println("hello world!", x, y)
	return "I have been returned!", 42
}

func increment(y *int) {
	*y = *y + 1
}

func sumArray(x [3]int) int {
	return x[0] + x[1] + x[2]
}

func sumArrayPointer(x *[3]int) int {
	return (*x)[0] + (*x)[1] + (*x)[2]
}

func sumArraySlice(x []int) int {
	return x[0] + x[1] + x[2]
}

func main() {
	fmt.Println(printHello("Mr", "Tay"))
	var x int
	x = 1
	increment(&x)
	fmt.Println(x)

	a := [3]int{1, 2, 3}
	fmt.Println(sumArray(a))

	b := [3]int{1, 2, 3}
	fmt.Println(sumArrayPointer(&b))

	//slices more efficient coz they are pointers
	//pass a slice argument
	c := []int{1,2,3}
	fmt.Println(sumArraySlice(c))
	
}
