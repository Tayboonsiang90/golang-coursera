package main

import "fmt"
// import "unicode"
// import "strings"

// UNUSED allows unused variables to be included in Go programs
func UNUSED(x ...interface{}) {}


func main() {
	// Single Line Comment...
	/* Block comment 1
	Block Comment 2 */

	//This prints a string
	fmt.Println("Hi")
	x := "Joe"
	fmt.Println("Hi "+ x)
	//formatting requires printf!!!
	fmt.Printf("Hi %s \n", x)

	//int8, int16, int32, int64
	//uint8, uint16, uint32, uint64
	//float32, float64
	//arithmetic + - * / % << >>
	//comparison == != > < >= <=
	//boolean && ||

	//type conversions
	var i int32 = 1
	var j int16 = 2
	// x = y //will fail
	i = int32(j) //T() operation
	UNUSED(i)

	var xf float64 = 123.45
	var yf float64 = 1.2345e2
	UNUSED(xf, yf)

	var z complex128 = complex(2,3)
	fmt.Println(z)
	
	//ASCII, each char is associated with 8 bit code.
	//'A' = 0x41
	//unicode is 32-bit
	//UTF-8 is variable length, subset of unicode, 8 bit to 32 bit. same as ASCII
	//Code point - unicode charcter
	//rune is a codepoint in go
	//unicode package test category of rune

	//unicode package
	//isdigit isspace isletter islower ispunct
	//toupper tolower

	//strings package
	//compare(a,b) 0 if equals, -1 <, 1 >
	//contains(s, substr)
	//hasprefix(s, prefix)
	//index(s, substr)
	//replace(s, old, new, n), returns new string
	//tolower(s), toupper(s)
	//trimspace(s) //leading and trailinh whitespace removed

	//strconv package
	//atoi(s)  string to int
	//itoa(s) int to string
	//formatfloat(f, fmt, prec, bitsize)
	//parsefloat(s, bitsize)

}

