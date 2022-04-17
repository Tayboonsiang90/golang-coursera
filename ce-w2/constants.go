package main

// UNUSED allows unused variables to be included in Go programs
func UNUSED(x ...interface{}) {}


func main() {
	const x = 1.3
	const (
		y = 4 
		z = "Hi"
	)

	type Grades int 
	const (
		A Grades = iota
		B
		C
		D
		F
	)
	
}