package main

import (
	"fmt"
	"net"
	"net/http"
	"encoding/json"
)

func main() {
	//net/http
	x, y := http.Get("www.uci.edu")
	fmt.Println(x, y)
	//net
	net.Dial("tcp", "uci.edu:80")

	type Person struct {
		name  string
		addr  string
		phone string
	}
	p1 := Person{name: "joe", addr: "bukit batok", phone: "123"}

	//json marshlling
	//[]byte, error=nil
	barr, _ := json.Marshal(p1)

	// p2 needs to fit the unmarshalled object
	var p2 Person
	//byte array, address of p2
	_ = json.Unmarshal(barr, &p2)
}
