package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	type Person struct {
		Name    string
		Address string
	}

	var p Person

	fmt.Println("Please enter a name: ")
	fmt.Scanln(&p.Name)
	fmt.Println("Please enter an address:")
	fmt.Scanln(&p.Address)

	fmt.Println(p)

	barr, _ := json.Marshal(p)

	fmt.Printf(string(barr))
}
