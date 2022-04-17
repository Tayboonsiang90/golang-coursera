package main

import (
	"fmt"
	"os"
)

func main() {

	type Person struct {
		FirstName string
		LastName  string
	}
	sli := make([]Person, 0, 0)

	var filename string
	fmt.Println("Please enter a filename: ")
	fmt.Scanln(&filename)
	f, _ := os.Open(filename)
	for {
		firstNameBArray := make([]byte, 20)
		spaceBArray := make([]byte, 1)
		lastNameBArray := make([]byte, 20)
		nextLineBArray := make([]byte, 2)
		_, err := f.Read(firstNameBArray)
		if err != nil {
			break
		}
		_, _ = f.Read(spaceBArray)
		_, _ = f.Read(lastNameBArray)
		_, _ = f.Read(nextLineBArray)

		var p Person
		p.FirstName = string(firstNameBArray)
		p.LastName = string(lastNameBArray)
		sli = append(sli, p)
	}
	f.Close()
	for _, v := range sli {
		fmt.Println(v)
	}
}
