package main

import (
	"fmt"
	"os"
)

func main() {
	// os.Open()
	// os.Close()
	// os.Read() //into []byte
	// os.Write() //with []byte

	//f is file handle, error handler
	f, _ := os.Open("dt.txt")
	barr := make([]byte, 10)
	//# of bytes read, err
	//reads and fills barr
	//read moves the head
	nb, _ := f.Read(barr)
	fmt.Println(nb)
	f.Close()

	f1, _ := os.Create("outfile1.txt")
	barr1 := []byte{1,2,3}
	nb1, _ := f1.Write(barr1)
	nb2, _ := f1.WriteString("lol")
	fmt.Println(nb1, nb2)
}
