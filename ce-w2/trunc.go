package main

import "fmt"

func main() {
	//Scan user input
	var x float32
	fmt.Println("Please enter a floating point number... ")
	fmt.Scan(&x)
	fmt.Printf("%d \n", int32(x))
}