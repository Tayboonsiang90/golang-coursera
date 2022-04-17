package main

import "fmt"

var funcVar func(int) int

func incFn(x int) int {
	return x + 1
}
func decFn(x int) int {
	return x - 1
}
func makeIncFn(amt int) func(int) int {
	return func(x int) int { return x + amt }
}

func applyIt(afunct func(int) int, val int) int {
	return afunct(val)
}

func fA() func() int {
	i := 5
	return func() int {
		i++
		return i
	}
}

//variable parameters treated as slice
//variadic functions
func getMax(vals ...int) int {
	lastV := -1
	for _, v := range vals {
		if v > lastV {
			lastV = v
		}
	}
	return lastV
}

func main() {
	//deferred function
	//cleanups
	defer fmt.Println("Bye!")

	funcVar = incFn
	fmt.Println(applyIt(funcVar, 1))
	funcVar = decFn
	fmt.Println(applyIt(funcVar, 1))

	//anonymous function
	fmt.Println(applyIt(
		func(x int) int {
			return x + 1
		}, 1))

	//return function
	fmt.Println(applyIt(makeIncFn(2), 2))

	//variadic function
	fmt.Println(getMax(1, 2, 3, 4, 5))
	fmt.Println(getMax([]int{1, 2, 3, 4, 5, 6}...)) //pass a slice in

	fB := fA()
	fmt.Print(fB())
	fmt.Print(fB())
}
