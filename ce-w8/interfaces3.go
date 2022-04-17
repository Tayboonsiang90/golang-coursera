package main


type Shape2D interface {
	Area() float64
	Perimeter() float64
}

type Triangle struct {
	base   float64
	height float64
}
func (t Triangle) Area() float64 {
	return (0.5 * t.base * t.height)
}
func (t Triangle) Perimeter() float64 {
	return (t.base + t.height)
}
type Rectangle struct {
	base   float64
	height float64
}
func (t Rectangle) Area() float64 {
	return (0.5 * t.base * t.height)
}
func (t Rectangle) Perimeter() float64 {
	return (t.base + t.height)
}

func FitInYard(s Shape2D) bool {
	if (s.Area() > 100 && s.Perimeter() > 100) {
		return true
	}
	return false
}


func main() {
	//All types satisfy an empty interface
	//use it to accept any type as parameter
	// func PrintMe(val interface{}) {
	// 	fmt.Println(val)
	// }
	
}

