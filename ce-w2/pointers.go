package main

import "fmt"

func main() {
	var x int = 1
	var y int
	var ip *int //points to an integer

	ip = &x //address to x
	fmt.Println(&x)
	y = *ip //value of the pointer
	fmt.Println(y)

	ptr := new(int) //default initialized to zero
	*ptr = 3 //value is now 3

	
}
