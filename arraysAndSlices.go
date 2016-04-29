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

	myNewSlice := make([]int, 2, 4)

	for i := 0; i < 20; i++ {
		myNewSlice = append(myNewSlice, i)
		fmt.Println("The capacity of my new Slice is: ", cap(myNewSlice))
	}
	fmt.Println(sliceOfSlice)
}
