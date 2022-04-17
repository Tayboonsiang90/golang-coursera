package main

import (
	"fmt"
	"time"
)

func prod(c chan int) {
	count := 0
	for {
		time.Sleep(1 * time.Second)
		c <- count
		count++
		if count == 15 {
			close(c)
		}
	}
}

func main() {
	c := make(chan int, 10)
	go prod(c)
	go prod(c)

	//need to close the channel when you are using range
	for i := range c {
		fmt.Println(i)
	}
}
