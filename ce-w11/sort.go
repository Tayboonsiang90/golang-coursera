package main

import (
	"fmt"
	"math"
	"sync"
	"sort"
)

//Bubble sort
func bbsort(sli []int, wg *sync.WaitGroup) {
	fmt.Println(sli)
	sort.Ints(sli)
	wg.Done()
}

func main() {
	arrayToBeSorted := [15]int{15, 14, 13, 12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	size := int(math.Floor(float64(len(arrayToBeSorted) / 4)))

	var wg sync.WaitGroup
	wg.Add(4)
	go bbsort(arrayToBeSorted[0:size], &wg)
	go bbsort(arrayToBeSorted[size:2*size], &wg)
	go bbsort(arrayToBeSorted[2*size:3*size], &wg)
	go bbsort(arrayToBeSorted[3*size:], &wg)
	wg.Wait()
}
