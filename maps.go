package main

import "fmt"

func main() {
	myMap := make(map[string]int)

	myMap["One"] = 1
	myMap["Two"] = 2
	myMap["Three"] = 3

	fmt.Println(myMap)

	for key, value := range myMap {
		fmt.Println(key, ": ", value)
	}

}
