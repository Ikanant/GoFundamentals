package main

import (
	"fmt"
)

func main() {
	//make will need (<type>, <len>, <cap>)

	mySlice := []int{1, 2, 3, 4, 5}
	sliceOfSlice := mySlice[0:3]

	fmt.Println(mySlice)
	fmt.Println(sliceOfSlice)

}
