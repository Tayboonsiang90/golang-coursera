package main

import (
	"fmt"
	"time"
)

var balance float64

//First subroutine prints available balance in the bank
func available() {
	balance++
	fmt.Printf("Balance: %f\n", balance)
}

//Second subroutine withdraws money from the bank
func withdraw(amount float64) {
	balance = balance + amount
	fmt.Printf("Balance: %f\n",balance)
}

func main() {
	go available()
	go withdraw(69)
	time.Sleep(1 * time.Second)

	//The race condition in this is if withdraw executes first, then available will print the wrong balance.
	//if withdraw executes second, then available will print the right balance.
}
