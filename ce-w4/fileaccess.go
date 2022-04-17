package main

import (
	"io/ioutil"
)

func main() {
	//open
	//read bytes into []byte
	//write []byte into file
	//close release
	//seek move read head

	//[]byte, error
	//large files will cause an error (eg. more than what ram is avilable)
	dat, _ := ioutil.ReadFile("test.txt")
	_ = dat

	x := make([]byte, 10)

	//err
	_ = ioutil.WriteFile("outfile.txt", x, 0777)

}
