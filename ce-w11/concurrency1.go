package main

import (
	"fmt" 
	"sync"
)

func foo(x *int, wg *sync.WaitGroup) {
	fmt.Println(*x)
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	x := 1
	wg.Add(1)
	go foo(&x, &wg)
	wg.Wait()
	x = x + 1
}
