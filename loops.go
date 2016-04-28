package main

import(
  "fmt"
)

func main(){

  //Infinite loop
  counter := 0
  for{
    if counter==5 {
      break
    } else {
      fmt.Print("Entered ", counter+1, " amount of times\n")
      counter++
    }
  }

  //Regular loop
  for i:=0; i<5; i++ {
    fmt.Println("Regular looping", i+1)
  }
}
