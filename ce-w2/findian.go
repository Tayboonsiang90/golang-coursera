package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	//Scan user input
	var x string
	fmt.Println("Please enter a string... ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	x = scanner.Text()
	if (strings.HasPrefix(strings.ToLower(x), "i") && strings.HasSuffix(strings.ToLower(x), "n") && strings.Contains(strings.ToLower(x), "a")) {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}