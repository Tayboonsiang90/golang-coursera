package main

import (
	"fmt"
	"time"
)

func main() {
	go fmt.Println("1")
	go fmt.Println("2")
	go fmt.Println("3")
	go fmt.Println("4")
	go fmt.Println("5")
	time.Sleep(2 * time.Second)
	fmt.Println("main")
	time.Sleep(200 * time.Millisecond)
}
