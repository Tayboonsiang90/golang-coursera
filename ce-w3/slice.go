package main

import (
	"fmt" 
	"strconv"
	"sort"
)

func main() {
	sli := make([]int, 0, 3)
	_ = sli
	for {
		fmt.Println("Please enter an integer: ")
		var s string
		fmt.Scan(&s)
		if (s=="X") {
			break
		}
		x, _ := strconv.Atoi(s)
		sli = append(sli, x)
		sort.Ints(sli)
		fmt.Println(sli)
	}
}
