package main

import (
	"fmt" // no more error
)

func main() {
	//exactly like objects
	var idMap map[string]int
	idMap = make(map[string]int)

	idMap = map[string]int{"joe": 123, "alice": 1234}

	fmt.Println(idMap["alice"])
	//will also change old key value pair
	idMap["jane"] = 123456
	fmt.Println(idMap["jane"])
	delete(idMap, "jane")
	fmt.Println(idMap["jane"]) //0

	//p boolean, id value (check if exists)
	id, p := idMap["joe"]
	fmt.Println(id, p)
	//how many key value pairs
	fmt.Println(len(idMap))

	//loop through
	for key, val := range idMap {
		fmt.Println(key, val)
	}
}
