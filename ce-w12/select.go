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
	c1 := make(chan int, 10)
	c2 := make(chan int, 10)
	go prod(c1)
	go prod(c2)
	for {
		//need to close the channel when you are using range
		select {
		case a := <-c1:
			fmt.Println(a)
		case b := <-c2:
			fmt.Println(b)
		}
	}
	// select {
	// case a := <-c1:
	// 	fmt.Println("recieved a")
	// case c2 <- b:
	// 	fmt.Println("sent b")
	// }
	// select {
	// case a := <-c1:
	// 	fmt.Println("recieved a")
	// case <- abort:
	// 	return
	// }
	// select {
	// case a := <-c1:
	// 	fmt.Println("recieved a")
	//  default:
	// 	fmt.Println("nothing")
	// }
}
