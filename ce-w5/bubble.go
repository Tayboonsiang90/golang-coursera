package main

import (
	"fmt"
	"strconv"
	"strings"
)

func BubbleSort(sliInt []int) {
	//Length of sliInt
	lenSli := len(sliInt)
	for i := 0; i < lenSli; i++ {
		for j := 0; j < lenSli-1; j++ {
			if sliInt[j] > sliInt[j+1] {
				Swap(sliInt, j)
			}
		}
	}
}

func Swap(sliInt []int, j int) {
	temp := sliInt[j]
	sliInt[j] = sliInt[j+1]
	sliInt[j+1] = temp
}

func main() {
	//Input and Input processing
	var seq string
	fmt.Println("type in a sequence of up to 10 integers, seperated by ,: ")
	fmt.Scanln(&seq)
	sliStr := strings.Split(seq, ",")
	sliInt := make([]int, 0, 0)
	for _, v := range sliStr {
		x, _ := strconv.Atoi(v)
		sliInt = append(sliInt, x)
	}

	//Bubble Sort
	BubbleSort(sliInt)

	//Display
	fmt.Println(sliInt)
}
