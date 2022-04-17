package main

import (
	"fmt"
	"math"
)

func GenDisplaceFn(acc, iVelo, iDisp float64) func(time float64) float64 {
	return func(time float64) float64 {
		return 0.5*acc*math.Pow(time, 2) + iVelo*time + iDisp
	}
}

func main() {
	var acc float64
	var iVelo float64
	var iDisp float64
	var time float64

	fmt.Println("Please enter acceleration: ")
	fmt.Scanln(&acc)
	fmt.Println("Please enter initial velocity: ")
	fmt.Scanln(&iVelo)
	fmt.Println("Please enter initial displacement: ")
	fmt.Scanln(&iDisp)
	fmt.Println("Please enter initial time: ")
	fmt.Scanln(&time)

	fn := GenDisplaceFn(acc, iVelo, iDisp)
	fmt.Println(fn(time))
}
