package main

import (
	"fmt"
	"sync"
)

var i int = 0
var mut sync.Mutex
var wg sync.WaitGroup

func inc() {
	mut.Lock()
	i = i + 1
	mut.Unlock()
	wg.Done()
}

func main() {
	//DONT LET TWO GOROUTINES WRITE AT THE SAME TIME
	//ACCESS TO SHARED VARIABLES NEED TO  HAVE MUTUAL EXCLUSION

	wg.Add(8)
	go inc()
	go inc()
	go inc()
	go inc()
	go inc()
	go inc()
	go inc()
	go inc()
	wg.Wait()
	fmt.Println(i)
}
