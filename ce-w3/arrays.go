package main

import (
	"fmt" // no more error
)

func main() {
	//ARRAYS
	var x [5]int
	x[0] = 2
	fmt.Printf("%d", x[1])

	var y = [5]int{1, 2, 3, 4, 5}
	var z = [...]int{1, 2, 3, 4, 5}

	for i, v := range x {
		fmt.Printf("index %d, value %d\n", i, v)
	}
	for i, v := range y {
		fmt.Printf("index %d, value %d\n", i, v)
	}
	for i, v := range z {
		fmt.Printf("index %d, value %d\n", i, v)
	}

	//SLICES
	//3 properties
	//1. Pointer to the start
	//2. Length of the slice
	//3. Capacity maximum of parent (depending on where the slice started)

	arr := [...]string{"a", "b", "c", "d", "e", "f", "g"}
	s1 := arr[1:3]
	s2 := arr[2:5]
	// length and capacity
	fmt.Println(len(s1), cap(s1))
	fmt.Println(len(s2), cap(s2))
	//using slice literals to reference underlying arrays
	sli := []int{1, 2, 3}
	_ = sli

	//3 ways to make a slice
	//1. make array, then call slice
	//2. use a slice literal
	//3. using make()
	//make (type, length/capacity)
	sli = make([]int, 10)
	//make (type, length, capacity)
	sli = make([]int, 10, 20)

	//append() adds elements to the end of the slice
	//increases the size of array if necessary
	sli = make([]int, 0, 3)
	fmt.Println(len(sli), cap(sli))
	sli = append(sli, 100)
	fmt.Println(len(sli), cap(sli))
}
