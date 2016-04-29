package main

import(
  "fmt"
)

func main(){
  //make will need (<type>, <len>, <cap>)

  mySlice := make([]string, 5, 10)
  fmt.Println("My slice length is: ", len(mySlice))
  fmt.Println("My slice capacity is: ", cap(mySlice))

}
