package main

import "fmt"

// UNUSED allows unused variables to be included in Go programs
func UNUSED(x ...interface{}) {}


func main() {
	var x int = 7
	if (x > 6) {
		fmt.Println(x)
	}
	for i := 0; i<5; i++ {
		fmt.Println(i)
	}
	// for {
	// 	fmt.Printf("Hi")
	// }
	//infinite loop

	switch x {
	case 7: 
		fmt.Printf("Case1")
	case 1:
		fmt.Printf("Case2")
	default: 
	fmt.Printf("default")
	}

	//tagless switch
	switch {
	case (x==7):
		fmt.Printf("Case 1")
	case (x==6):
		fmt.Printf("Case 2")}

	//break
	//continue skips the iteration

	//Scan user input
	var y *int
	fmt.Println(y) //nil
	var i int
	y = &i
	fmt.Printf("Number of y?")
	num, err := fmt.Scan(y)
	fmt.Println(*y)
	UNUSED(num, err)
}
